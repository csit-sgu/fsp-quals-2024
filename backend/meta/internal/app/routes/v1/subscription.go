package v1

import (
	"fmt"
	"net/http"

	"app/internal/app/appcontext"
	"app/internal/app/ckey"
	"app/internal/app/errors"
	"app/internal/config"
	"app/internal/log"
	"app/internal/mail"
	"app/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Subscription controller
//
//	@summary	Leave an email subscription request
//	@tags		Subscription
//	@accept		json
//
//	@param		request	body	model.Subscription	true	"Subscription Info"
//	@produce	json
//
//	@router		/subscription [POST]
func PostSubscription(c *gin.Context) {
	traceId := c.GetString(string(ckey.TraceId))
	ctx := c.Request.Context()
	l := log.L().TraceId(traceId)

	var body model.Subscription
	if err := c.ShouldBind(&body); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		c.Abort()
		return
	}

	found, err := appcontext.Ctx.Clickhouse.FindSubByEmail(ctx, body.Email)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		c.Abort()
		return
	}
	if found {
		_ = c.Error(errors.E().Code(errors.CodeBadInput).TraceId(traceId).Message("A subscription for this email already exists").Build()).
			SetType(gin.ErrorTypePublic)
		c.Abort()
		return
	}

	uid, err := uuid.NewV7()
	if err != nil {
		log.S.Error(
			"Failed to generate subscription confirmation token",
			l.Error(err),
		)
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		c.Abort()
		return
	}
	body.Confirmation = uid

	mailBody := fmt.Sprintf(config.C.Mail.Body, body.Confirmation.String())
	err = mail.SendEmail(body.Email, config.C.Mail.Subject, mailBody)
	if err != nil {
		log.S.Info("Failed to send an email", l.Error(err))
		_ = c.Error(errors.E().Code(errors.CodeBadInput).TraceId(traceId).Message("Failed to send a confirmation email").Build()).
			SetType(gin.ErrorTypePublic)
		c.Abort()
		return
	}

	if err := appcontext.Ctx.Clickhouse.SaveSubscription(ctx, body); err != nil {
		log.S.Error("Failed to submit subscription", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		c.Abort()
		return
	}

	c.Status(http.StatusCreated)
}

// Subscription confirmation controller
//
//	@summary	Confirm an email subscription request
//	@tags		Subscription
//	@accept		json
//
//	@param		request	body	model.SubscriptionConfirmation	true	"Subscription Confirmation"
//	@produce	json
//
//	@success	201	{string}	string	"Confirmed"
//
//	@router		/subscription/confirm [POST]
func ConfirmSubscription(c *gin.Context) {
	traceId := c.GetString(string(ckey.TraceId))
	ctx := c.Request.Context()

	var body model.SubscriptionConfirmation
	if err := c.ShouldBind(&body); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		c.Abort()
		return
	}

	found, err := appcontext.Ctx.Clickhouse.FindSubscription(
		c,
		body.Confirmation,
	)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		c.Abort()
		return
	}

	if found {
		if err := appcontext.Ctx.Clickhouse.ActivateSubscription(ctx, body.Confirmation); err != nil {
			_ = c.Error(err).SetType(gin.ErrorTypePublic)
			c.Abort()
			return
		}
	} else {
		_ = c.Error(errors.E().Code(errors.CodeBadInput).Message("Subscription not found").TraceId(traceId).Build()).SetType(gin.ErrorTypePublic)
		c.Abort()
		return
	}

	c.Status(http.StatusCreated)
}

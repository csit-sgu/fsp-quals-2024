package v1

import (
	"net/http"

	"app/internal/app/appcontext"
	"app/internal/app/ckey"
	"app/internal/log"
	"app/internal/model"

	"github.com/gin-gonic/gin"
)

// Subscription controller
//
//	@summary	Leave an email subscription request
//	@tags		Subscription
//	@accept		json
//
//	@param		message	body	model.Subscription	true	"Subscription Info"
//	@produce	json
//
//	@success	201	{string}	string	"Created"
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

	if err := appcontext.Ctx.Clickhouse.SaveSubscription(ctx, body); err != nil {
		log.S.Error("Failed to submit subscription", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.Status(http.StatusCreated)
}

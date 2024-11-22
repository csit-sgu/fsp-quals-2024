package v1

import (
	"net/http"

	"app/internal/app/appcontext"
	"app/internal/log"
	"app/internal/model"

	"github.com/gin-gonic/gin"
)

// Filter controller.
//
//	@summary	Filter event data
//	@tags		Filter
//	@accept		plain
//
//	@produce	json
//	@param		request	body		model.FilterRequest	true	"Query params"
//
//	@success	200		{string}	model.Event
//
//	@router		/filter [post]
func FilterData(c *gin.Context) {
	traceId := c.GetString("trace_id")
	ctx := c.Request.Context()
	l := log.L().TraceId(traceId)
	var resp []model.Event
	defer c.JSON(http.StatusOK, &resp)

	var r model.FilterRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		log.S.Error("Failed to bind query", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	log.S.Debug("Query parameters validation successful", l)

	resp, err := appcontext.Ctx.Clickhouse.FilterEvents(
		ctx,
		l,
		r.Condition,
		r.RequiredFields,
	)
	if err != nil {
		log.S.Error("Failed to filter events", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
}

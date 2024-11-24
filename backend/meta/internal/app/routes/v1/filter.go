package v1

import (
	"net/http"

	"app/internal/app/appcontext"
	"app/internal/db/opensearch"
	"app/internal/log"
	"app/internal/model"

	"github.com/gin-gonic/gin"
)

// Filter controller.
//
//	@summary	Filter event data
//	@tags		Filter
//	@accept		application/json
//
//	@produce	json
//	@param		request	body		model.FilterRequest	true	"Filter data"
//
//	@success	200		{string}	model.FilterResponse
//
//	@router		/filter [post]
func FilterData(c *gin.Context) {
	traceId := c.GetString("trace_id")
	ctx := c.Request.Context()
	l := log.L().TraceId(traceId)
	var resp model.FilterResponse

	var r model.FilterRequest
	if err := c.ShouldBind(&r); err != nil {
		log.S.Error("Failed to bind query", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	log.S.Debug("Query parameters validation successful", l.Add("cond", r))

	resp, err := appcontext.Ctx.Clickhouse.FilterEvents(
		ctx,
		l,
		r,
	)

	resp.Events, err = opensearch.ApplyFuzzySearch(
		ctx,
		l,
		r.Condition.Title,
		r.Condition.AdditionalInfo,
		resp.Events,
	)
	if err != nil {
		log.S.Error("Failed to filter events", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	log.S.Debug("Successfully filtered events", l)

	c.JSON(http.StatusOK, &resp)
}

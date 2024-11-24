package v1

import (
	"app/internal/app/appcontext"
	"app/internal/log"
	"app/internal/model"

	"github.com/gin-gonic/gin"
)

// Notify controller.
//
//	@summary	Notify service with about changes in database
//	@tags
//	@accept		json
//
//	@produce	json
//
//	@success	200		{string}	[]string
//	@param		request	body		model.NotifyRequest	true	"Updated events codes"
//
//	@router		/notify [POST]
func Notify(c *gin.Context) {
	traceId := c.GetString("trace_id")
	l := log.L().TraceId(traceId)

	r := model.NotifyRequest{}
	if err := c.ShouldBind(&r); err != nil {
		log.S.Error("Failed to bind query", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	log.S.Debug("Query parameters were validated successful", l.Add("cond", r))

	data, err := appcontext.Ctx.Clickhouse.GetIndexData(c, l)

	err = appcontext.Ctx.OpenSearch.IndexData(c, l, data)
	if err != nil {
		log.S.Error("Failed to index data", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	if err != nil {
		log.S.Error("Failed to filter events", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
}

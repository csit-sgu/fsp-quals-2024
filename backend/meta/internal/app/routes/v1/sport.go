package v1

import (
	"net/http"

	"app/internal/app/appcontext"
	"app/internal/log"

	"github.com/gin-gonic/gin"
)

// Sports controller.
//
//	@summary	Get available sports
//	@tags		Sport
//	@accept		plain
//
//	@produce	json
//
//	@success	200	{string}	[]string
//
//	@router		/sports [GET]
func GetSports(c *gin.Context) {
	traceId := c.GetString("trace_id")
	ctx := c.Request.Context()
	l := log.L().TraceId(traceId)
	var resp []string
	defer c.JSON(http.StatusOK, &resp)

	resp, err := appcontext.Ctx.Clickhouse.GetSports(ctx)
	if err != nil {
		log.S.Error("Failed to filter events", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
}

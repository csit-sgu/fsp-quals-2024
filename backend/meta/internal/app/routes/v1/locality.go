package v1

import (
	"net/http"

	"app/internal/app/appcontext"
	"app/internal/log"

	"github.com/gin-gonic/gin"
)

// Country controller.
//
//	@summary	Get available countries
//	@tags		Location
//	@accept		plain
//
//	@produce	json
//
//	@success	200	{string}	[]string
//
//	@router		/countries [GET]
func GetCountry(c *gin.Context) {
	traceId := c.GetString("trace_id")
	ctx := c.Request.Context()
	l := log.L().TraceId(traceId)
	var resp []string
	defer c.JSON(http.StatusOK, &resp)

	resp, err := appcontext.Ctx.Clickhouse.GetCountries(ctx)
	if err != nil {
		log.S.Error("Failed to filter events", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
}

// Region controller.
//
//	@summary	Get available regions
//	@tags		Location
//	@accept		plain
//
//	@produce	json
//	@param		country	query		string	true	"Query params"
//
//	@success	200	{string}	[]string
//
//	@router		/regions [GET]
func GetRegion(c *gin.Context) {
	traceId := c.GetString("trace_id")
	ctx := c.Request.Context()
	l := log.L().TraceId(traceId)
	var resp []string
	defer c.JSON(http.StatusOK, &resp)

	country := c.Query("country")

	log.S.Debug(
		"Query parameters validation successful",
		l.Add("country", country),
	)

	resp, err := appcontext.Ctx.Clickhouse.GetRegions(ctx, country)
	if err != nil {
		log.S.Error("Failed to filter events", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
}

// Locality controller.
//
//	@summary	Get available localities
//	@tags		Location
//	@accept		plain
//
//	@produce	json
//	@param		country	query		string	true	"Country query parameter"
//	@param		region	query		string	true	"Region query parameter"
//
//	@success	200	{string}	[]string
//
//	@router		/localities [GET]
func GetLocalities(c *gin.Context) {
	traceId := c.GetString("trace_id")
	ctx := c.Request.Context()
	l := log.L().TraceId(traceId)
	var resp []string
	defer c.JSON(http.StatusOK, &resp)

	country := c.Query("country")
	region := c.Query("region")

	log.S.Debug(
		"Query parameters validation successful",
		l.Add("country", country).Add("region", region),
	)

	resp, err := appcontext.Ctx.Clickhouse.GetLocalities(ctx, country, region)
	if err != nil {
		log.S.Error("Failed to fetch localities", l.Error(err))
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
}

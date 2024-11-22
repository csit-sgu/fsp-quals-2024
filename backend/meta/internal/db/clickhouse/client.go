package clickhouse

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"app/internal/config"
	"app/internal/log"
	"app/internal/model"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type ClickhouseClient struct {
	conn driver.Conn
}

func InitClickhouseClient(c config.DatabaseConfig) *ClickhouseClient {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{c.Host + ":" + c.Port},
		Auth: clickhouse.Auth{
			Database: c.Database,
			Username: c.Username,
			Password: c.Password,
		},
		Debug: true,
	})
	if err != nil {
		log.S.Error("Failed to connect to ClickHouse", log.L().Error(err))
	}

	return &ClickhouseClient{conn}
}

func (c ClickhouseClient) extractWhereParts(
	cond model.FilterCondition,
) (parts []string, namedFields []any) {
	t := reflect.TypeOf(cond)
	v := reflect.ValueOf(cond)
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("db")
		if !v.Field(i).IsZero() {
			parts = append(parts, fmt.Sprintf("%s = @%s", key, key))
			namedFields = append(namedFields, driver.NamedValue{
				Name:  key,
				Value: v.Field(i),
			})
		}
	}

	return parts, namedFields
}

func (c ClickhouseClient) BuildFilterQuery(
	cond model.FilterCondition,
	fields []string,
) (query string, namedFields []any) {
	var fieldPart string
	if len(fields) == 0 {
		fieldPart = "*"
	} else {
		fieldPart = strings.Join(fields, ",")
	}
	whereClause := ""

	parts, namedFields := c.extractWhereParts(cond)

	whereParts := strings.Join(parts, " AND ")

	if whereParts != "" {
		whereClause = "WHERE " + whereParts
	} else {
		whereClause = ""
	}

	query = "SELECT " + fieldPart + " " + "FROM db.events" + " " + whereClause + ";"
	// log.S.Debug("Built filter query", log.L().Add("query", query))

	return query, namedFields
}

func (c ClickhouseClient) FilterEvents(
	ctx context.Context,
	l log.LogObject,
	cond model.FilterCondition,
	fields []string,
) (events []model.Event, err error) {
	query, namedFields := c.BuildFilterQuery(cond, fields)

	if err = c.conn.Select(ctx, &events, query, namedFields...); err != nil {
		log.S.Error(
			"Failed to execute filter query",
			log.L().Add("query", query).Add("error", err),
		)
		return nil, err
	} else {
		log.S.Debug("Events were retrived successfully", l.Add("count", len(events)))
	}

	return events, nil
}

func (c *ClickhouseClient) GetCountries(
	ctx context.Context,
) (countries []string, err error) {
	var rows []model.Country
	if err := c.conn.Select(ctx, &rows, countryQuery); err != nil {
		log.S.Error(
			"Failed to execute country query",
			log.L().Add("query", countryQuery).Add("error", err),
		)
		return nil, err
	}

	for i := range rows {
		countries = append(countries, rows[i].Country)
	}

	log.S.Debug(
		"Countries were retrived successfully",
		log.L().Add("count", len(countries)),
	)

	return countries, nil
}

func (c *ClickhouseClient) GetRegions(
	ctx context.Context,
	country string,
) (regions []string, err error) {
	var rows []model.Region
	if err := c.conn.Select(ctx, &rows, regionQuery, clickhouse.Named("country", country)); err != nil {
		log.S.Error(
			"Failed to execute region query",
			log.L().Add("query", countryQuery).Add("error", err),
		)
		return nil, err
	}

	log.S.Debug(
		"Regions were retrived successfully",
		log.L().Add("count", len(regions)),
	)

	for i := range rows {
		regions = append(regions, rows[i].Region)
	}

	return regions, nil
}

func (c *ClickhouseClient) GetLocalities(
	ctx context.Context,
	country string,
	region string,
) (localities []string, err error) {
	var rows []model.Locality
	if err := c.conn.Select(ctx, &rows, localityQuery, clickhouse.Named("country", country), clickhouse.Named("region", region)); err != nil {
		log.S.Error(
			"Failed to execute locality query",
			log.L().Add("query", localityQuery).Add("error", err),
		)
		return nil, err
	}

	log.S.Debug(
		"Localities were retrieved successfully",
		log.L().Add("count", len(rows)),
	)

	for i := range rows {
		localities = append(localities, rows[i].Locality)
	}

	return localities, nil
}

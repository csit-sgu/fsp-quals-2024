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
	fieldPart := strings.Join(fields, ",")
	whereClause := ""

	parts, namedFields := c.extractWhereParts(cond)

	whereParts := strings.Join(parts, " AND ")

	if whereParts != "" {
		whereClause = "WHERE " + whereParts
	} else {
		whereClause = ""
	}

	query = "SELECT " + fieldPart + " " + "FROM events" + " " + whereClause + ";"
	// log.S.Debug("Built filter query", log.L().Add("query", query))

	return query, namedFields
}

func (c ClickhouseClient) FilterEvents(
	ctx context.Context,
	cond model.FilterCondition,
	fields []string,
) (events []model.Event, err error) {
	query, namedFields := c.BuildFilterQuery(cond, fields)

	if c.conn.Select(ctx, &events, query, namedFields...); err != nil {
		log.S.Error(
			"Failed to execute filter query",
			log.L().Add("query", query).Add("error", err),
		)
		return nil, err
	}

	return events, nil
}

func (c *ClickhouseClient) GetCountries(
	ctx context.Context,
) (countries []string, err error) {
	results := []string{}
	if err := c.conn.Select(ctx, &countries, countryQuery); err != nil {
		log.S.Error(
			"Failed to execute country query",
			log.L().Add("query", countryQuery).Add("error", err),
		)
		return nil, err
	}

	return results, nil
}

func (c *ClickhouseClient) GetRegions(
	ctx context.Context,
	country string,
) (regions []string, err error) {
	results := []string{}
	if err := c.conn.Select(ctx, &regions, countryQuery, country); err != nil {
		log.S.Error(
			"Failed to execute country query",
			log.L().Add("query", countryQuery).Add("error", err),
		)
		return nil, err
	}

	return results, nil
}

func (c *ClickhouseClient) GetLocalities(
	ctx context.Context,
	region string,
) (regions []string, err error) {
	results := []string{}
	if err := c.conn.Select(ctx, &regions, countryQuery, region); err != nil {
		log.S.Error(
			"Failed to execute country query",
			log.L().Add("query", countryQuery).Add("error", err),
		)
		return nil, err
	}

	return results, nil
}

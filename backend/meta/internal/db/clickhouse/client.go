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

func (c ClickhouseClient) buildCondition(
	key string,
	fieldValue reflect.Value,
	fieldType string,
) (string, []any) {
	switch fieldType {
	case "common":
		return fmt.Sprintf(
				"%s = @%s",
				key,
				key,
			), []any{
				driver.NamedValue{Name: key, Value: fieldValue},
			}
	case "interval":
		return fmt.Sprintf(
				"%s BETWEEN @%s_from AND @%s_to",
				key,
				key,
				key,
			), []any{
				driver.NamedValue{
					Name:  key + "_from",
					Value: fieldValue.FieldByName("From").Interface(),
				},
				driver.NamedValue{
					Name:  key + "_to",
					Value: fieldValue.FieldByName("To").Interface(),
				},
			}
	default:
		return "", nil
	}
}

func (c ClickhouseClient) extractWhereParts(
	cond model.FilterCondition,
) (parts []string, namedFields []any) {
	t := reflect.TypeOf(cond)
	v := reflect.ValueOf(cond)
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("ch")
		if !v.Field(i).IsZero() {
			fieldType := t.Field(i).Tag.Get("filter")
			part, newFields := c.buildCondition(key, v.Field(i), fieldType)
			if part == "" {
				continue
			}
			log.S.Debug("Condition part", log.L().Add("part", part))
			if part != "" {
				parts = append(parts, part)
				namedFields = append(namedFields, newFields...)
			}
		}
	}

	log.S.Debug("Condition parts", log.L().Add("parts", parts))

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
	log.S.Debug("Built filter query", log.L().Add("query", query))

	return query, namedFields
}

func (c ClickhouseClient) FilterEvents(
	ctx context.Context,
	l log.LogObject,
	cond model.FilterCondition,
	fields []string,
) (events []model.Event, err error) {
	query, namedFields := c.BuildFilterQuery(cond, fields)

	var eventViews []model.EventView
	if err = c.conn.Select(ctx, &eventViews, query, namedFields...); err != nil {
		log.S.Error(
			"Failed to execute filter query",
			log.L().Add("query", query).Add("error", err),
		)
		return nil, err
	} else {
		log.S.Debug("Events were retrived successfully", l.Add("count", len(events)))
	}

	for i := range eventViews {
		locationData, err := c.getLocationData(ctx, eventViews[i].Code)
		if err != nil {
			log.S.Warn("Failed to get location data", log.L().Add("error", err))
			return nil, err
		}
		ageData, err := c.getAgeData(ctx, eventViews[i].Code)
		if err != nil {
			log.S.Warn("Failed to get age data", log.L().Add("error", err))
			return nil, err
		}
		events = append(events, model.Event{
			Code:           eventViews[i].Code,
			StartDate:      eventViews[i].StartDate,
			LocationData:   locationData,
			AgeData:        ageData,
			Title:          eventViews[i].Title,
			AdditionalInfo: eventViews[i].AdditionalInfo,
			Participants:   eventViews[i].Participants,
			Stage:          eventViews[i].Stage,
			EndDate:        eventViews[i].EndDate,
			Sport:          eventViews[i].Sport,
		})
	}

	return events, nil
}

func (c ClickhouseClient) getAgeData(
	ctx context.Context,
	code string,
) (ageData []model.AgeData, err error) {
	if err := c.conn.Select(ctx, &ageData, ageQuery, clickhouse.Named("code", code)); err != nil {
		log.S.Error(
			"Failed to execute age query",
			log.L().Add("query", ageQuery).Add("error", err),
		)
		return nil, err
	}

	return ageData, nil
}

func (c ClickhouseClient) getLocationData(
	ctx context.Context,
	code string,
) (locationData []model.LocationData, err error) {
	if err := c.conn.Select(ctx, &locationData, locationQuery, clickhouse.Named("code", code)); err != nil {
		log.S.Error(
			"Failed to execute age query",
			log.L().Add("query", locationQuery).Add("error", err),
		)
		return nil, err
	}

	return locationData, nil
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

func (c *ClickhouseClient) GetSports(
	ctx context.Context,
) (sports []string, err error) {
	var rows []model.Sport
	if err := c.conn.Select(ctx, &rows, sportsQuery); err != nil {
		log.S.Error(
			"Failed to execute region query",
			log.L().Add("query", countryQuery).Add("error", err),
		)
		return nil, err
	}

	log.S.Debug(
		"Regions were retrived successfully",
		log.L().Add("count", len(sports)),
	)

	for i := range rows {
		sports = append(sports, rows[i].Sport)
	}

	return sports, nil
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

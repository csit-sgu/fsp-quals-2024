package clickhouse

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"app/internal/log"
	"app/internal/model"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

var viewFields = [...]string{
	"code",
	"start_date",
	"country",
	"region",
	"locality",
	"gender",
	"left_bound",
	"right_bound",
	"title",
	"additional_info",
	"n_participants",
	"stage",
	"end_date",
	"sport",
	"extra_mapping",
	"page_index",
}

func (c ClickhouseClient) buildCommonCondition(
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
			part, newFields := c.buildCommonCondition(
				key,
				v.Field(i),
				fieldType,
			)
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

	if !reflect.ValueOf(cond.Age).IsZero() {
		part, nameField := c.getAgeRestrictionQuery(cond.Age, cond.Gender)
		parts = append(parts, part)
		namedFields = append(namedFields, nameField)
	}

	log.S.Debug("Condition parts", log.L().Add("parts", parts))

	return parts, namedFields
}

func (c ClickhouseClient) getAgeRestrictionQuery(
	age uint32,
	gender string,
) (query string, namedField any) {
	return "db.general_view.left_bound <= @age AND db.general_view.right_bound >= @age",
		clickhouse.Named(
			"age",
			age,
		)
}

func (c ClickhouseClient) BuildFilterQuery(
	request model.FilterRequest,
) (query string, namedFields []any, whereClause string) {
	var fieldPart string
	fields := request.RequiredFields
	cond := request.Condition
	pagination := request.Pagination
	if len(fields) == 0 {
		fieldPart = strings.Join(viewFields[:], ",")
	} else {
		fieldPart = strings.Join(fields, ",")
	}
	whereClause = ""

	paginationPart := "o.page_index >= @page_lower AND o.page_index <= @page_upper"

	parts, namedFields := c.extractWhereParts(cond)
	namedFields = append(
		namedFields,
		clickhouse.Named("page_lower", pagination.Page*pagination.PageSize),
	)
	namedFields = append(
		namedFields,
		clickhouse.Named("page_upper", (pagination.Page+1)*pagination.PageSize),
	)

	whereParts := strings.Join(parts, " AND ")

	if whereParts != "" {
		whereClause = "WHERE " + whereParts
	} else {
		whereClause = ""
	}

	query = fmt.Sprintf(filterQuery, whereClause, fieldPart, paginationPart)

	log.S.Debug("Built filter query", log.L().Add("query", query))

	return query, namedFields, whereClause
}

func (c ClickhouseClient) GetIndexData(
	ctx context.Context,
	l log.LogObject,
	request model.NotifyRequest,
) (indexData []model.IndexData, err error) {
	query := codeQuery
	arg := clickhouse.Named("codes", &request)

	if err = c.conn.Select(ctx, &indexData, query, arg); err != nil {
		log.S.Error(
			"Failed to execute index query",
			log.L().Add("query", query).Error(err),
		)
		return nil, err
	} else {
		log.S.Debug(
			"Index data were retrived successfully",
			log.L().Add("count", len(indexData)),
		)
	}

	return indexData, nil
}


type Count struct {
	Count uint64 `ch:"count"`
}

func (c ClickhouseClient) FilterEvents(
	ctx context.Context,
	l log.LogObject,
	request model.FilterRequest,
) (response model.FilterResponse, err error) {
	query, namedFields, whereClause := c.BuildFilterQuery(request)

	mapping := make(map[string]model.Event)
	filterView := model.FilterView{}
	total := []Count{}
	eventViews := filterView.Events
	if err = c.conn.Select(ctx, &eventViews, query, namedFields...); err != nil {
		log.S.Error(
			"Failed to execute filter query",
			log.L().Add("query", query).Add("error", err),
		)
		return model.FilterResponse{}, err
	} else {
		log.S.Debug("Events were retrieved successfully", l.Add("count", len(eventViews)))
	}
	if err = c.conn.Select(ctx, &total, fmt.Sprintf(
		filterCounterQuery,
		whereClause,
	), namedFields...); err != nil {
		log.S.Error(
			"Failed to execute filter counter query",
			log.L().Add("query", filterCounterQuery).Add("error", err),
		)
		return model.FilterResponse{}, err
	}
	filterView.Total = uint32(total[0].Count)

	for i := range eventViews {
		currentEvent, ok := mapping[eventViews[i].Code]
		view := eventViews[i]

		locData := model.LocationData{
			Country:  view.Country,
			Region:   view.Region,
			Locality: view.Locality,
		}
		ageData := model.AgeData{
			Gender:     view.Gender,
			LeftBound:  view.LeftBound,
			RightBound: view.RightBound,
			Original:   view.ExtraMapping,
		}
		if ok {
			currentEvent.LocationData = append(
				currentEvent.LocationData,
				locData,
			)

			currentEvent.AgeData = append(
				currentEvent.AgeData,
				ageData,
			)
		} else {
			event := model.Event{
				Code:           view.Code,
				StartDate:      model.CustomTime(view.StartDate),
				LocationData:   []model.LocationData{locData},
				AgeData:        []model.AgeData{ageData},
				Title:          view.Title,
				AdditionalInfo: view.AdditionalInfo,
				Participants:   view.Participants,
				Stage:          view.Stage,
				EndDate:        model.CustomTime(view.EndDate),
				Sport:          view.Sport,
			}

			response.Events = append(response.Events, &event)

			mapping[view.Code] = event
		}
	}

	response.Total = total[0].Count

	return response, nil
}

package clickhouse

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"app/internal/log"
	"app/internal/model"

	"github.com/ClickHouse/clickhouse-go/v2"
)

var viewFields = [...]string{
	"o.code",
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
	"end_date",
	"sport",
	"extra_mapping",
	"page_index",
    "event_type",
    "event_scale",
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

func (c ClickhouseClient) buildPart(
	structField reflect.StructField,
	fieldValue reflect.Value,
) (part string, namedFields []any) {
	chTag := structField.Tag.Get("ch")
	filterTag := structField.Tag.Get("filter")

	switch filterTag {
	case "common":
		return fmt.Sprintf(
				"%s = @%s",
				chTag,
				chTag,
			), []any{
				clickhouse.Named(chTag,fieldValue.Interface()),
			}
	case "interval":
		return fmt.Sprintf(
				"%s BETWEEN @%s_from AND @%s_to",
				chTag,
				chTag,
				chTag,
			), []any{
				clickhouse.Named(chTag+"_from", fieldValue.FieldByName("From").Interface()),
				clickhouse.Named(chTag+"_to", fieldValue.FieldByName("To").Interface()),
			}
	case "inside":
		return fmt.Sprintf(
				"@%s >= left_bound AND @%s <= right_bound",
				chTag,
				chTag,
			), []any{
                clickhouse.Named(chTag, fieldValue.Interface()),
			}
	}
	return part, namedFields
}

func (c ClickhouseClient) buildWhereClause(
	cond model.FilterCondition,
	fieldNames []string,
) (whereClause string, namedFields []any) {
	var parts []string
	for _, fieldName := range fieldNames {

		structField, ok := reflect.TypeOf(cond).FieldByName(fieldName)
		log.S.Debug("Struct field", log.L().Add("field", structField))
		if !ok {
			panic("This should not ever happen")
		}
		// get field value
		fieldValue := reflect.ValueOf(cond).FieldByName(fieldName)

		if !fieldValue.IsZero() {
			part, newFields := c.buildPart(structField, fieldValue)

			parts = append(parts, part)
			namedFields = append(namedFields, newFields...)
		}

		log.S.Debug("Field parts", log.L().Add("parts", parts))
		log.S.Debug("Field named fields", log.L().Add("namedFields", namedFields))
	}

	if len(parts) > 0 {
		whereClause = "WHERE "
		log.S.Debug("Where clause", log.L().Add("whereClause", parts))
	} else {
		whereClause = ""
	}
	whereClause += strings.Join(parts, " AND ")
	return whereClause, namedFields
}

var commonFields = []string{
	"Code",
	"EventScale",
	"EventType",
	"Sport",
}

var locationFields = []string{
	"Country",
	"Region",
	"Locality",
}

var ageFields = []string{
	"Age",
	"Gender",
}

func (c ClickhouseClient) BuildFilterQuery(
	request model.FilterRequest,
) (query string, countQuery string, namedFields []any) {
	pagination := request.Pagination
	selectFields := request.RequiredFields
	cond := request.Condition
	var selectPart string
	if len(selectFields) == 0 {
		selectPart = strings.Join(viewFields[:], ",")
	} else {
		selectPart = strings.Join(selectFields, ",")
	}

	commonWhere, commonNamedFields := c.buildWhereClause(cond, commonFields)
	locationWhere, locationNamedFields := c.buildWhereClause(cond, locationFields)
	ageWhere, ageNamedFields := c.buildWhereClause(cond, ageFields)

	namedFields = append(
		namedFields,
		commonNamedFields...,
	)
	namedFields = append(
		namedFields,
		locationNamedFields...,
	)
	namedFields = append(
		namedFields,
		ageNamedFields...,
	)
	namedFields = append(
		namedFields,
		clickhouse.Named("page_lower", pagination.Page*pagination.PageSize),
		clickhouse.Named("page_upper", (pagination.Page+1)*pagination.PageSize),
	)

	paginationPart := "WHERE page_index >= @page_lower AND page_index <= @page_upper"

	query = fmt.Sprintf(
		filterQuery,
		locationWhere,
		ageWhere,
		commonWhere,
		selectPart,
		paginationPart,
	)
	log.S.Debug("Built filter query", log.L().Add("query", query).Add("namedFields", namedFields))

	countQuery = fmt.Sprintf(
		filterCounterQuery,
		locationWhere,
		ageWhere,
		commonWhere,
	)

	return query, countQuery, namedFields
}

func (c ClickhouseClient) GetIndexData(
	ctx context.Context,
	l log.LogObject,
	request model.NotifyRequest,
) (indexData []model.IndexData, err error) {
	query := codeQuery

	if err = c.conn.Select(ctx, &indexData, query); err != nil {
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
	query, countQuery, namedFields := c.BuildFilterQuery(request)

	mapping := make(map[string]model.Event)
	filterView := model.FilterView{}
	total := []Count{}
	eventViews := filterView.Events
	if err = c.conn.Select(ctx, &eventViews, query, namedFields...); err != nil {
		log.S.Error(
			"Failed to execute filter query",
			log.L().Add("query", query).Error(err),
		)
		return model.FilterResponse{}, err
	} else {
		log.S.Debug("Events were retrieved successfully", l.Add("count", len(eventViews)))
	}
	if err = c.conn.Select(ctx, &total, countQuery, namedFields...); err != nil {
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

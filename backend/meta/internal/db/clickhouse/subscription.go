package clickhouse

import (
	"context"

	"app/internal/log"
	"app/internal/model"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func (c *ClickhouseClient) SaveSubscription(
	ctx context.Context,
	sub model.Subscription,
) error {
	err := c.conn.Exec(
		ctx,
		subInsertQuery,
		clickhouse.Named("email", sub.Filter.Code),
		clickhouse.Named("code", sub.Filter.Code),
		clickhouse.Named("gender", sub.Filter.Gender),
		clickhouse.Named("age", sub.Filter.Age),
		clickhouse.Named("sport", sub.Filter.Sport),
		clickhouse.Named("additional_info", sub.Filter.AdditionalInfo),
		clickhouse.Named("country", sub.Filter.Country),
		clickhouse.Named("region", sub.Filter.Region),
		clickhouse.Named("locality", sub.Filter.Locality),
		clickhouse.Named("event_type", sub.Filter.EventType),
		clickhouse.Named("event_scale", sub.Filter.EventScale),
		clickhouse.Named("start_date", sub.Filter.DateRange.From),
		clickhouse.Named("end_date", sub.Filter.DateRange.To))
	if err != nil {
		log.S.Error(
			"Failed to save subscription request",
			log.L().Add("query", subInsertQuery).Add("error", err),
		)
		return err
	}

	log.S.Debug("Subscription was succefully saved", log.L())
	return nil
}

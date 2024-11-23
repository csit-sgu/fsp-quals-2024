package clickhouse

import (
	"context"

	"app/internal/log"
	"app/internal/model"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/google/uuid"
)

func (c *ClickhouseClient) FindSubByEmail(
	ctx context.Context,
	email string,
) (bool, error) {
	var sub []model.Subscription
	if err := c.conn.Select(ctx, &sub, subFindByMail, clickhouse.Named("email", email)); err != nil {
		log.S.Error(
			"Subscription find query has failed",
			log.L().Add("query", subFindByMail).Add("error", err),
		)
		return false, err
	}

	if len(sub) == 0 {
		log.S.Info(
			"Didn't find subscription by email in the database",
			log.L().Add("query", subFindByMail),
		)
		return false, nil
	}

	log.S.Debug("Subscription was succefully found", log.L())
	return true, nil
}

func (c *ClickhouseClient) SaveSubscription(
	ctx context.Context,
	sub model.Subscription,
) error {
	err := c.conn.Exec(
		ctx,
		subInsertQuery,
		clickhouse.Named("email", sub.Code),
		clickhouse.Named("code", sub.Code),
		clickhouse.Named("confirmation", sub.Confirmation),
		clickhouse.Named("gender", sub.Gender),
		clickhouse.Named("age", sub.Age),
		clickhouse.Named("sport", sub.Sport),
		clickhouse.Named("additional_info", sub.AdditionalInfo),
		clickhouse.Named("country", sub.Country),
		clickhouse.Named("region", sub.Region),
		clickhouse.Named("locality", sub.Locality),
		clickhouse.Named("event_type", sub.EventType),
		clickhouse.Named("event_scale", sub.EventScale),
		clickhouse.Named("start_date", sub.StartDate),
		clickhouse.Named("end_date", sub.EndDate))
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

func (c *ClickhouseClient) FindSubscription(
	ctx context.Context,
	id uuid.UUID,
) (bool, error) {
	var count uint64
	if err := c.conn.QueryRow(ctx, subCountQuery, clickhouse.Named("confirmation", id)).Scan(&count); err != nil {
		log.S.Error(
			"Subscription find query has failed",
			log.L().Add("query", subCountQuery).Add("error", err),
		)
		return false, err
	}

	if count == 0 {
		log.S.Info(
			"Didn't find subscription confirmation in the database",
			log.L().Add("query", subCountQuery),
		)
		return false, nil
	}

	log.S.Debug("Subscription was succefully found", log.L())
	return true, nil
}

func (c *ClickhouseClient) ActivateSubscription(
	ctx context.Context,
	id uuid.UUID,
) error {
	if err := c.conn.Exec(ctx, subActivateQuery, clickhouse.Named("confirmation", id)); err != nil {
		log.S.Error(
			"Failed to activate subscription",
			log.L().Add("query", subActivateQuery).Add("error", err),
		)
		return err
	}

	log.S.Debug("Subscription was succefully activated", log.L())
	return nil
}

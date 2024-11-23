package clickhouse

import (
	"context"

	"app/internal/log"
	"app/internal/model"

	"github.com/ClickHouse/clickhouse-go/v2"
)

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

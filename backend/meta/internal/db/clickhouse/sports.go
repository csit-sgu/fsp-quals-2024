package clickhouse

import (
	"context"

	"app/internal/log"
	"app/internal/model"
)

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

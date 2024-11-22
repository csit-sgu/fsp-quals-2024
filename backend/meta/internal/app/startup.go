package app

import (
	"context"

	"app/internal/config"
	"app/internal/db/clickhouse"
)

// Add all required onShutdown logic here.
func onStartup(ctx context.Context) error {
	client := clickhouse.InitClickhouseClient(config.C.Database.Clickhouse)

	Ctx = &AppContext{
		Clickhouse: client,
	}
	return nil
}
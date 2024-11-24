package app

import (
	"context"

	"app/internal/app/appcontext"
	"app/internal/config"
	"app/internal/db/clickhouse"
	"app/internal/log"
)

// Add all required onShutdown logic here.
func onStartup(ctx context.Context) error {
	chClient, err := clickhouse.InitClickhouseClient(
		config.C.Database.Clickhouse,
	)
	if err != nil {
		return err
	}

	appcontext.Ctx = &appcontext.AppContext{
		Clickhouse: chClient,
	}

	log.S.Debug("App startup is complete", log.L())

	return nil
}

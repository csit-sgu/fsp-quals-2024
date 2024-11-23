package app

import (
	"context"

	"app/internal/app/appcontext"
	"app/internal/config"
	"app/internal/db/clickhouse"
	"app/internal/db/opensearch"
)

// Add all required onShutdown logic here.
func onStartup(ctx context.Context) error {
	chClient, err := clickhouse.InitClickhouseClient(config.C.Database.Clickhouse)

	if err != nil {
		return err
	}

	osClient, err := opensearch.InitOpenSearchClient(config.C.Database.OpenSearch)

	if err != nil {
		return err
	}

	appcontext.Ctx = &appcontext.AppContext{
		Clickhouse: chClient,
		OpenSearch: osClient,
	}
	return nil
}

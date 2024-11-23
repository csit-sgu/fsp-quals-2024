package app

import (
	"context"
	"strings"

	"app/internal/app/appcontext"
	"app/internal/config"
	"app/internal/db/clickhouse"
	"app/internal/db/os"
	"app/internal/log"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

// Add all required onShutdown logic here.
func onStartup(ctx context.Context) error {
	chClient, err := clickhouse.InitClickhouseClient(
		config.C.Database.Clickhouse,
	)
	if err != nil {
		return err
	}

	osClient, err := os.InitOpenSearchClient(config.C.Database.OpenSearch)
	if err != nil {
		return err
	}

	settings := strings.NewReader(`{
        "settings": {
            "number_of_shards": 1,
            "number_of_replicas": 0
        }
    }`)

	_ = opensearchapi.IndicesCreateRequest{
		Index: config.C.Database.OpenSearch.Index,
		Body:  settings,
	}

    log.S.Debug("OpenSearch index was created", log.L().Add("index", config.C.Database.OpenSearch.Index))

	appcontext.Ctx = &appcontext.AppContext{
		Clickhouse: chClient,
		OpenSearch: &os.OpenSearch{
			Client: osClient,
			Index: config.C.Database.OpenSearch.Index,
		},
	}

    log.S.Debug("App startup is complete", log.L())

	return nil
}

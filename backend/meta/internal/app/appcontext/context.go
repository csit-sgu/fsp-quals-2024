package appcontext

import (
	"app/internal/db/clickhouse"

	"github.com/opensearch-project/opensearch-go"
)

type AppContext struct {
	Clickhouse *clickhouse.ClickhouseClient
    OpenSearch *opensearch.Client
}

var Ctx *AppContext

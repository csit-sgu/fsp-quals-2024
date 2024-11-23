package appcontext

import (
	"app/internal/db/clickhouse"
	"app/internal/db/os"
)

type AppContext struct {
	Clickhouse *clickhouse.ClickhouseClient
	OpenSearch *os.OpenSearch
}

var Ctx *AppContext

package appcontext

import (
	"app/internal/db/clickhouse"
)

type AppContext struct {
	Clickhouse *clickhouse.ClickhouseClient
}

var Ctx *AppContext

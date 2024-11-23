package clickhouse

import (
	"app/internal/config"
	"app/internal/log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type ClickhouseClient struct {
	conn driver.Conn
}

func InitClickhouseClient(c config.DatabaseConfig) *ClickhouseClient {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{c.Host + ":" + c.Port},
		Auth: clickhouse.Auth{
			Database: c.Database,
			Username: c.Username,
			Password: c.Password,
		},
		Debug: true,
	})
	if err != nil {
		log.S.Error("Failed to connect to ClickHouse", log.L().Error(err))
	}

	return &ClickhouseClient{conn}
}

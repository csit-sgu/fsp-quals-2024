package sql

import (
	"fmt"

	"app/internal/config"

	"github.com/jmoiron/sqlx"
)

var C *Connection

type Connection struct {
	db *sqlx.DB
}

// Connects to the database using credentials provided in the config.
func Connect(config config.DatabaseConfig) (c *Connection, err error) {
	var db *sqlx.DB
	if db, err = sqlx.Connect(
		config.Database,
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
            config.Database,
		),
	); err != nil {
		return nil, err
	}

	c = &Connection{
		db: db,
	}
	return c, nil
}

// Closes database connection.
func (c *Connection) Close() {
	c.db.Close()
}

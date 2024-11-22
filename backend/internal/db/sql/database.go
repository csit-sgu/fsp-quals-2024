package sql

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"app/internal/config"
	"app/internal/log"
	"app/internal/model"

	"github.com/jmoiron/sqlx"
)

var C *Connection

type Connection struct {
	db *sqlx.DB
}

// Connects to the database using credentials provided in the config.
func Connect() (c *Connection, err error) {
	var db *sqlx.DB
	if db, err = sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			config.C.Database.Username,
			config.C.Database.Password,
			config.C.Database.Host,
			config.C.Database.Port,
			config.C.Database.Name,
		),
	); err != nil {
		return
	}

	c = &Connection{
		db: db,
	}
	return
}

func (c *Connection) extractWhereParts(
	cond model.FilterCondition,
) (parts []string) {
	t := reflect.TypeOf(cond)
	v := reflect.ValueOf(cond)
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("db")
		if !v.Field(i).IsZero() {
			parts = append(parts, fmt.Sprintf("%s = :%s", key, key))
		}
	}

	return parts
}

func (c *Connection) BuildFilterQuery(
	cond model.FilterCondition,
	fields []string,
) (query string) {
	fieldPart := strings.Join(fields, ",")
	whereClause := ""
	whereParts := strings.Join(c.extractWhereParts(cond), " AND ")

	if whereParts != "" {
		whereClause = "WHERE " + whereParts
	} else {
		whereClause = ""
	}

	query = "SELECT " + fieldPart + " " + "FROM events" + " " + whereClause
	// log.S.Debug("Built filter query", log.L().Add("query", query))

	return query
}

func (c *Connection) FilterEvents(
	ctx context.Context,
	cond model.FilterCondition,
	fields []string,
) (sql.Result, error) {
	query := c.BuildFilterQuery(cond, fields)

	rows, err := c.db.NamedExec(query, &cond)
	if err != nil {
		log.S.Error(
			"Failed to execute filter query",
			log.L().Add("query", query).Add("error", err),
		)
		return nil, err
	}

	return rows, nil
}

// Closes database connection.
func (c *Connection) Close() {
	c.db.Close()
}

// +build go1.10

package driver

import (
	"database/sql"
)

func OpenDB(dsn, name string) *sql.DB {
	tc := &mockConnector{
		dsn:      dsn,
		instance: name,
	}
	return sql.OpenDB(tc)
}

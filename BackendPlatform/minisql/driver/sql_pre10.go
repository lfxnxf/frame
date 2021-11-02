// +build !go1.10

package driver

import "database/sql"

func OpenDB(dsn, name string) *sql.DB {
	d, _ := sql.Open("mock-mysql", dsn)
	return d
}

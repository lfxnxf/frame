package driver

import (
	"context"
	"database/sql/driver"

	"github.com/go-sql-driver/mysql"
)

var mysqlDriver = &mysql.MySQLDriver{}

type mockConnector struct {
	dsn      string
	instance string
}

func (t *mockConnector) Connect(c context.Context) (driver.Conn, error) {
	conn, err := mysqlDriver.Open(t.dsn)
	if err != nil {
		return nil, err
	}
	return &mockConn{conn}, err
}

func (t *mockConnector) Driver() driver.Driver {
	return mysqlDriver
}

type MockDriver struct {
	Dsn string
}

func (m *MockDriver) Open(name string) (driver.Conn, error) {
	conn, err := mysqlDriver.Open(m.Dsn)
	if err != nil {
		return nil, err
	}
	return &mockConn{conn}, err
}

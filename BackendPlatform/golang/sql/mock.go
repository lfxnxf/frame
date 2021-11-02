package sql

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/BackendPlatform/minisql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

// NewMock returns sqlmock.Sqlmock and db.close() func.
// sqlmock.Sqlmock doc shows https://godoc.org/github.com/DATA-DOG/go-sqlmock
func NewMock(name string) (mock sqlmock.Sqlmock, closeFunc func(), err error) {
	group, mock, closeFunc, err := NewMockGroup()
	SQLGroupManager.Add(name, group)
	return
}

// NewMockGroup returns builtin mock group and sqlmock.
// sqlmock.Sqlmock can mock data for all SQL command
func NewMockGroup() (group *Group, mock sqlmock.Sqlmock, closeFunc func(), err error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logging.Errorf("init sqlmock err,err(%v)", err)
	}
	closeFunc = func() {
		db.Close()
	}
	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		logging.Errorf("open gorm err,err(%v)", err)
		return
	}
	client := &Client{DB: gormDB}
	group = &Group{
		master:  client,
		replica: []*Client{client},
		next:    0,
		total:   1,
	}
	return
}

// NewMockSQL returns builtin mock group and close function.
// Optional sqlPath be used to init mysql,eg: create table or insert data.
func NewMockSQL(sqlPath ...string) (group *Group, closeFunc func(), err error) {
	var buf []byte
	s := minisql.Run()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true&loc=Local",
		s.User, s.Password, s.Host, s.Port, s.DB)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logging.Warnf("sql open err,err: %v", err)
		return
	}
	closeFunc = func() {
		db.Close()
		s.Stop()
	}
	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		logging.Errorf("open gorm err,err(%v)", err)
		return
	}
	client := &Client{DB: gormDB}
	group = &Group{
		master:  client,
		replica: []*Client{client},
		next:    0,
		total:   1,
	}
	if len(sqlPath) > 0 {
		buf, err = ioutil.ReadFile(sqlPath[0])
		if err != nil {
			logging.Warnf("read sql file err,err: %v", err)
			return
		}
		err = gormDB.Exec(string(buf)).Error
	}
	return
}

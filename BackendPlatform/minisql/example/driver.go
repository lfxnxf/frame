package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lfxnxf/frame/BackendPlatform/minisql/driver"
)

func main() {
	db, err := sql.Open("mock-mysql", "")
	if err != nil {
		return
	}
	err = db.Ping()
	fmt.Println(err)
	db.Close()

}

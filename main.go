package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:password@tcp(object_store)/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Printf("successfully pinged db with dsn: %v\n", dsn)
}

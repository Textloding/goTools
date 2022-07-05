package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	sql.Register("mysql", &MySQLDriver{})
}

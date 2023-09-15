package configs

import (
	"database/sql"
	"fmt"
	"os"
)

func NewMysql() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	return sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, dbName))
}

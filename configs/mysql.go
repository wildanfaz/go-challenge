package configs

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func NewMySQL() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbPort := os.Getenv("MYSQL_PORT")
	dbHost := os.Getenv("MYSQL_HOST")

	sql, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, dbHost, dbPort, dbName))

	if err != nil {
		return nil, err
	}

	sql.SetConnMaxIdleTime(15 * time.Minute)
	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(50)

	return sql, nil
}

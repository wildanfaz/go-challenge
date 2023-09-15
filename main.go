package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/wildanfaz/go-challenge/cmd"
)

func init() {

}

func main() {
	if err := cmd.Start(); err != nil {
		panic(err)
	}
}

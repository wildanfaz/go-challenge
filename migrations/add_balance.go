package migrations

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/wildanfaz/go-challenge/configs"
)

func AddBalance(email string) {
	db, err := configs.NewMySQL()

	if err != nil {
		panic(err)
	}

	res, err := db.ExecContext(context.Background(), `
	UPDATE users SET balance = balance + 1000000 WHERE email=?
	`, email)

	af, err := res.RowsAffected()

	if err != nil {
		log.Error(err)
		return
	}

	if af < 1 {
		log.Warn("no rows affected")
		return
	}

	if err != nil {
		log.Error(err)
		return
	}

	log.Info("add balance success")
}

package migrations

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/wildanfaz/go-challenge/configs"
)

func Dumy() {
	db, err := configs.NewMySQL()

	if err != nil {
		panic(err)
	}

	_, err = db.ExecContext(context.Background(), `
		INSERT INTO products(name, description, category, price) VALUES
		('Melon Juice', 'Fresh', 'Drink', 10000),
		('Avocado Juice', 'Fresh', 'Drink', 10000),
		('Fried Rice', 'Delicious', 'Food', 12000),
		('Fried Banana', 'Delicious', 'Food', 2000)
		`)

	if err != nil {
		log.Error(err)
		return
	}

	log.Info("insert dumy success")
}

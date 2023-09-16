package migrations

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/wildanfaz/go-challenge/configs"
)

func Migrate() {
	db, err := configs.NewMySQL()

	if err != nil {
		panic(err)
	}

	_, err = db.ExecContext(context.Background(), `
	CREATE TABLE test.users (
		id varchar(36) DEFAULT (uuid()) NOT NULL,
		email varchar(255) UNIQUE NOT NULL,
		password varchar(255) NOT NULL,
		balance int NOT NULL DEFAULT 0,
		created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
		updated_at TIMESTAMP DEFAULT current_timestamp ON UPDATE current_timestamp NOT NULL,
		CONSTRAINT users_PK PRIMARY KEY (id)
	)
	`)

	if err != nil {
		log.Error(err)
		return
	}

	_, err = db.ExecContext(context.Background(), `
	CREATE TABLE test.products (
		id varchar(36) DEFAULT (uuid()) NOT NULL,
		name varchar(255) NOT NULL,
		description text,
		category varchar(255) NOT NULL,
		price int NOT NULL DEFAULT 0,
		created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
		updated_at TIMESTAMP DEFAULT current_timestamp ON UPDATE current_timestamp NOT NULL,
		CONSTRAINT products_PK PRIMARY KEY (id)
	)
	`)

	if err != nil {
		log.Error(err)
		return
	}

	_, err = db.ExecContext(context.Background(), `
	CREATE TABLE test.carts (
		id varchar(36) DEFAULT (uuid()) NOT NULL,
		user_id varchar(255) NOT NULL REFERENCES users(id),
		product_id varchar(255) NOT NULL REFERENCES products(id),
		amount int NOT NULL DEFAULT 1,
		created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
		updated_at TIMESTAMP DEFAULT current_timestamp ON UPDATE current_timestamp NOT NULL,
		CONSTRAINT carts_PK PRIMARY KEY (id)
	)
	`)

	if err != nil {
		log.Error(err)
		return
	}

	_, err = db.ExecContext(context.Background(), `
	CREATE TABLE test.payments (
		id varchar(36) DEFAULT (uuid()) NOT NULL,
		cart_id varchar(255) NOT NULL REFERENCES carts(id),
		created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
		updated_at TIMESTAMP DEFAULT current_timestamp ON UPDATE current_timestamp NOT NULL,
		CONSTRAINT carts_PK PRIMARY KEY (id)
	)
	`)

	if err != nil {
		log.Error(err)
		return
	}

	log.Info("migrate success")
}

func Rollback() {
	db, err := configs.NewMySQL()

	if err != nil {
		panic(err)
	}

	_, err = db.ExecContext(context.Background(), `
	DROP TABLE IF EXISTS users
	`)

	if err != nil {
		log.Error(err)
		return
	}

	_, err = db.ExecContext(context.Background(), `
	DROP TABLE IF EXISTS products
	`)

	if err != nil {
		log.Error(err)
		return
	}

	_, err = db.ExecContext(context.Background(), `
	DROP TABLE IF EXISTS carts
	`)

	if err != nil {
		log.Error(err)
		return
	}

	_, err = db.ExecContext(context.Background(), `
	DROP TABLE IF EXISTS payments
	`)

	if err != nil {
		log.Error(err)
		return
	}

	log.Info("rollback success")
}

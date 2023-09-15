package routers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-challenge/configs"
	"github.com/wildanfaz/go-challenge/internal/app/middlewares"
	"github.com/wildanfaz/go-challenge/internal/app/repositories"
	"github.com/wildanfaz/go-challenge/internal/app/services/products"
	"github.com/wildanfaz/go-challenge/internal/app/services/users"
)

func New() {
	app := fiber.New()

	// init db
	db, err := configs.NewMysql()

	if err != nil {
		panic(err)
	}

	// init repo
	Auth := repositories.NewAuthRepo(db)
	Products := repositories.NewProductsRepo(db)

	// init service
	register := users.NewRegister(Auth)
	login := users.NewLogin(Auth)
	addToShoppingCart := products.NewAddToShoppingCart(Products)
	checkoutAndPayment := products.NewCheckoutAndPayment(Products)
	deleteInShoppingCart := products.NewDeleteInShoppingCart(Products)
	listInShoppingCart := products.NewListInShoppingCart(Products)
	listProducts := products.NewListProducts(Products)

	app.Post("/register", register)
	app.Post("/login", login)

	// add middleware
	app.Use(middlewares.Auth)

	app.Post("/add/cart", addToShoppingCart)
	app.Post("/checkout", checkoutAndPayment)
	app.Delete("/delete/cart/:id", deleteInShoppingCart)
	app.Get("/list/cart", listInShoppingCart)
	app.Get("/list/products", listProducts)

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = ":3000"
	}

	log.Fatal(app.Listen(port))
}

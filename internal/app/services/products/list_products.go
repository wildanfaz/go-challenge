package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-challenge/internal/app/repositories"
	"github.com/wildanfaz/go-challenge/internal/app/services"
)

type list_products struct {
	repo repositories.Products
}

func NewListProducts(repo repositories.Products) services.Service {
	var listProducts = list_products{repo: repo}

	return listProducts.Service
}

func (str *list_products) Service(c *fiber.Ctx) error {
	panic(1)
}

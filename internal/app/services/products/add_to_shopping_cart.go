package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-challenge/internal/app/repositories"
	"github.com/wildanfaz/go-challenge/internal/app/services"
)

type add_to_shopping_cart struct {
	repo repositories.Products
}

func NewAddToShoppingCart(repo repositories.Products) services.Service {
	var addToShoppingCart = add_to_shopping_cart{repo: repo}

	return addToShoppingCart.Service
}

func (str *add_to_shopping_cart) Service(c *fiber.Ctx) error {
	return c.JSON("OKE")
}

package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-challenge/internal/app/repositories"
	"github.com/wildanfaz/go-challenge/internal/app/services"
)

type list_in_shopping_cart struct {
	repo repositories.Products
}

func NewListInShoppingCart(repo repositories.Products) services.Service {
	var listInShoppingCart = list_in_shopping_cart{repo: repo}

	return listInShoppingCart.Service
}

func (str *list_in_shopping_cart) Service(c *fiber.Ctx) error {
	panic(1)
}

package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-challenge/internal/app/repositories"
	"github.com/wildanfaz/go-challenge/internal/app/services"
)

type delete_in_shopping_cart struct {
	repo repositories.Products
}

func NewDeleteInShoppingCart(repo repositories.Products) services.Service {
	var deleteInShoppingCart = delete_in_shopping_cart{repo: repo}

	return deleteInShoppingCart.Service
}

func (str *delete_in_shopping_cart) Service(c *fiber.Ctx) error {
	panic(1)
}

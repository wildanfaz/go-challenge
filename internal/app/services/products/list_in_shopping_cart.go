package products

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/wildanfaz/go-challenge/internal/app/helpers"
	"github.com/wildanfaz/go-challenge/internal/app/repositories"
	"github.com/wildanfaz/go-challenge/internal/app/services"
	"github.com/wildanfaz/go-challenge/internal/app/types"
)

type list_in_shopping_cart struct {
	repo repositories.Products
}

func NewListInShoppingCart(repo repositories.Products) services.Service {
	var listInShoppingCart = list_in_shopping_cart{repo: repo}

	return listInShoppingCart.Service
}

func (str *list_in_shopping_cart) Service(c *fiber.Ctx) error {
	products, err := str.repo.ListProductsInCart(c.Context())

	if err != nil {
		log.Errorf("list products in cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	log.Info(types.ListProductsInCart)
	return helpers.NewResponse(c, http.StatusOK, types.ListProductsInCart, nil, products)
}

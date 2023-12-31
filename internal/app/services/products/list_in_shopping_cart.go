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
	user repositories.Auth
}

func NewListInShoppingCart(repo repositories.Products, user repositories.Auth) services.Service {
	var listInShoppingCart = list_in_shopping_cart{repo: repo, user: user}

	return listInShoppingCart.Service
}

func (str *list_in_shopping_cart) Service(c *fiber.Ctx) error {
	user, err := str.user.GetUserByEmail(c.Context(), c.Context().Value("email").(string))

	if err != nil {
		log.Errorf("list products in cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	products, err := str.repo.ListProductsInCart(c.Context(), user.ID.String())

	if err != nil {
		log.Errorf("list products in cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	log.Info(types.ListProductsInCart)
	return helpers.NewResponse(c, http.StatusOK, types.ListProductsInCart, nil, products)
}

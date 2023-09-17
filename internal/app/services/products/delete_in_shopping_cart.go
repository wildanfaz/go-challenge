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

type delete_in_shopping_cart struct {
	repo repositories.Products
	user repositories.Auth
}

func NewDeleteInShoppingCart(repo repositories.Products, user repositories.Auth) services.Service {
	var deleteInShoppingCart = delete_in_shopping_cart{repo: repo, user: user}

	return deleteInShoppingCart.Service
}

func (str *delete_in_shopping_cart) Service(c *fiber.Ctx) error {
	var (
		productID = c.Params("product_id")
	)

	user, err := str.user.GetUserByEmail(c.Context(), c.Context().Value("email").(string))

	if err != nil {
		log.Errorf("delete product in cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	exist, err := str.repo.CheckProduct(c.Context(), productID)

	if err != nil {
		log.Errorf("delete product in cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	if !exist {
		log.Warn(types.ErrProductNotFound)
		return helpers.NewResponse(c, http.StatusNotFound, types.Default, types.ErrProductNotFound, nil)
	}

	exist, err = str.repo.CheckProductInCart(c.Context(), user.ID.String(), productID)

	if err != nil {
		log.Errorf("delete product in cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	if !exist {
		log.Warn(types.ErrProductInCartNotFound)
		return helpers.NewResponse(c, http.StatusNotFound, types.Default, types.ErrProductInCartNotFound, nil)
	}

	if err := str.repo.DeleteInShoppingCart(c.Context(), user.ID.String(), productID); err != nil {
		log.Errorf("delete product in cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	log.Info(types.DeleteProductInCart)
	return helpers.NewResponse(c, http.StatusOK, types.DeleteProductInCart, nil, nil)
}

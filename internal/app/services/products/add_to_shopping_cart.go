package products

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/wildanfaz/go-challenge/internal/app/entities"
	"github.com/wildanfaz/go-challenge/internal/app/helpers"
	"github.com/wildanfaz/go-challenge/internal/app/repositories"
	"github.com/wildanfaz/go-challenge/internal/app/services"
	"github.com/wildanfaz/go-challenge/internal/app/types"
)

type add_to_shopping_cart struct {
	repo repositories.Products
	user repositories.Auth
}

func NewAddToShoppingCart(repo repositories.Products, user repositories.Auth) services.Service {
	var addToShoppingCart = add_to_shopping_cart{repo: repo, user: user}

	return addToShoppingCart.Service
}

func (str *add_to_shopping_cart) Service(c *fiber.Ctx) error {
	var (
		addProductToCart entities.AddProductToCart
	)

	if err := c.BodyParser(&addProductToCart); err != nil {
		log.Errorf("add product to cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusUnprocessableEntity, types.Default, types.ErrRequestBody, nil)
	}

	user, err := str.user.GetUserByEmail(c.Context(), c.Context().Value("email").(string))

	if err != nil {
		log.Errorf("add product to cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	addProductToCart.UserID = user.ID.String()

	if err := validation.ValidateStruct(&addProductToCart,
		validation.Field(&addProductToCart.UserID, validation.Required),
		validation.Field(&addProductToCart.ProductID, validation.Required),
		validation.Field(&addProductToCart.Amount, validation.Required, validation.Min(1)),
	); err != nil {
		log.Errorf("add product to cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusUnprocessableEntity, types.Default, err, nil)
	}

	exist, err := str.repo.CheckProduct(c.Context(), addProductToCart.ProductID)

	if err != nil {
		log.Errorf("add product to cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	if !exist {
		log.Warn(types.ErrProductNotFound)
		return helpers.NewResponse(c, http.StatusNotFound, types.Default, types.ErrProductNotFound, nil)
	}

	if err := str.repo.AddProductToCart(c.Context(), addProductToCart); err != nil {
		log.Errorf("add product to cart got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	log.Info(types.ProductToCart)
	return helpers.NewResponse(c, http.StatusCreated, types.ProductToCart, nil, nil)
}

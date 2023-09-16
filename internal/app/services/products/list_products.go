package products

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/wildanfaz/go-challenge/internal/app/entities"
	"github.com/wildanfaz/go-challenge/internal/app/helpers"
	"github.com/wildanfaz/go-challenge/internal/app/repositories"
	"github.com/wildanfaz/go-challenge/internal/app/services"
	"github.com/wildanfaz/go-challenge/internal/app/types"
)

type list_products struct {
	repo repositories.Products
}

func NewListProducts(repo repositories.Products) services.Service {
	var listProducts = list_products{repo: repo}

	return listProducts.Service
}

func (str *list_products) Service(c *fiber.Ctx) error {
	var (
		product entities.Product
	)

	product.Category = c.Query("category")

	products, err := str.repo.ListProducts(c.Context(), product)

	if err != nil {
		log.Errorf("list products got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	log.Info(types.ListProducts)
	return helpers.NewResponse(c, http.StatusOK, types.ListProducts, nil, products)
}

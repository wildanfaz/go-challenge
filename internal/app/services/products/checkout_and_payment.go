package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-challenge/internal/app/repositories"
	"github.com/wildanfaz/go-challenge/internal/app/services"
)

type checkout_and_payment struct {
	repo repositories.Products
}

func NewCheckoutAndPayment(repo repositories.Products) services.Service {
	var checkoutAndPayment = checkout_and_payment{repo: repo}

	return checkoutAndPayment.Service
}

func (str *checkout_and_payment) Service(c *fiber.Ctx) error {
	panic(1)
}

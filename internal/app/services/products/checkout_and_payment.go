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

type checkout_and_payment struct {
	repo repositories.Products
	user repositories.Auth
}

func NewCheckoutAndPayment(repo repositories.Products, user repositories.Auth) services.Service {
	var checkoutAndPayment = checkout_and_payment{repo: repo, user: user}

	return checkoutAndPayment.Service
}

func (str *checkout_and_payment) Service(c *fiber.Ctx) error {
	user, err := str.user.GetUserByEmail(c.Context(), c.Context().Value("email").(string))

	if err != nil {
		log.Errorf("checkout and payment got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	balance, enough, err := str.repo.CheckBalance(c.Context(), user.ID.String())

	if err != nil {
		log.Errorf("checkout and payment got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	if !enough {
		log.Warn(types.ErrInsufficientBalance)
		return helpers.NewResponse(c, http.StatusUnprocessableEntity, types.Default, types.ErrInsufficientBalance, nil)
	}

	err = str.repo.CheckoutAndPayment(c.Context(), balance, user.ID.String())

	if err != nil && err != types.ErrCartNotFound {
		log.Errorf("checkout and payment got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	if err != nil && err == types.ErrCartNotFound {
		log.Errorf("checkout and payment got error : %v", err)
		return helpers.NewResponse(c, http.StatusNotFound, types.Default, types.ErrCartNotFound, nil)
	}

	log.Info(types.CheckoutAndPayment)
	return helpers.NewResponse(c, http.StatusCreated, types.CheckoutAndPayment, nil, nil)
}

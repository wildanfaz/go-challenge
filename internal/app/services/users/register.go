package users

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/wildanfaz/go-challenge/internal/app/entities"
	"github.com/wildanfaz/go-challenge/internal/app/helpers"
	"github.com/wildanfaz/go-challenge/internal/app/repositories"
	"github.com/wildanfaz/go-challenge/internal/app/services"
	"github.com/wildanfaz/go-challenge/internal/app/types"
	"golang.org/x/crypto/bcrypt"
)

type register struct {
	repo repositories.Auth
}

func NewRegister(repo repositories.Auth) services.Service {
	var register = register{repo: repo}

	return register.Service
}

func (str *register) Service(c *fiber.Ctx) error {
	var (
		user entities.User
	)

	if err := c.BodyParser(&user); err != nil {
		log.Errorf("register got error : %v", err)
		return helpers.NewResponse(c, http.StatusUnprocessableEntity, types.Default, types.ErrRequestBody, nil)
	}

	if err := validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 0)),
	); err != nil {
		log.Errorf("register got error : %v", err)
		return helpers.NewResponse(c, http.StatusUnprocessableEntity, types.Default, err, nil)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Errorf("register got error : %v", err)
		return helpers.NewResponse(c, http.StatusUnprocessableEntity, types.Default, types.ErrHash, nil)
	}

	user.Password = string(hashed)

	if err := str.repo.Register(c.Context(), user); err != nil {
		log.Errorf("register got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	log.Info(types.Register)
	return helpers.NewResponse(c, http.StatusCreated, types.Register, nil, nil)
}

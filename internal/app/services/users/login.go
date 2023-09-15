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
	"github.com/wildanfaz/go-challenge/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	repo repositories.Auth
}

func NewLogin(repo repositories.Auth) services.Service {
	var login = login{repo: repo}

	return login.Service
}

func (str *login) Service(c *fiber.Ctx) error {
	var (
		user entities.User
	)

	if err := c.BodyParser(&user); err != nil {
		log.Errorf("login got error : %v", err)
		return helpers.NewResponse(c, http.StatusUnprocessableEntity, types.Default, types.ErrRequestBody, nil)
	}

	if err := validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 0)),
	); err != nil {
		log.Errorf("login got error : %v", err)
		return helpers.NewResponse(c, http.StatusUnprocessableEntity, types.Default, err, nil)
	}

	userDb, err := str.repo.GetUserByEmail(c.Context(), user.Email)

	if err != nil {
		log.Errorf("login got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrDatabase, nil)
	}

	if userDb.Email == "" {
		log.Warn("user not found")
		return helpers.NewResponse(c, http.StatusNotFound, types.Default, types.ErrUserNotFound, nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password)); err != nil {
		log.Errorf("login got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrComparePassword, nil)
	}

	ss, err := auth.GenerateToken(user.Email)

	if err != nil {
		log.Errorf("login got error : %v", err)
		return helpers.NewResponse(c, http.StatusInternalServerError, types.Default, types.ErrInvalidToken, nil)
	}

	return helpers.NewResponse(c, http.StatusOK, types.Login, nil, map[string]string{
		"token": ss,
	})
}

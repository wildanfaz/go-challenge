package middlewares

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/wildanfaz/go-challenge/internal/app/helpers"
	"github.com/wildanfaz/go-challenge/internal/app/types"
	"github.com/wildanfaz/go-challenge/pkg/auth"
)

func Auth(c *fiber.Ctx) error {
	header := c.GetReqHeaders()["Authorization"]

	bearerToken := strings.Split(header, " ")

	if len(bearerToken) < 2 {
		log.Info("invalid bearer token")
		return helpers.NewResponse(c, http.StatusUnauthorized, types.Default, types.ErrInvalidToken, nil)
	}

	claims, err := auth.ValidateToken(bearerToken[1])

	if err != nil {
		log.Error("middleware got error : %v", err)
		return helpers.NewResponse(c, http.StatusUnauthorized, types.Default, types.ErrInvalidToken, nil)
	}

	c.Context().SetUserValue("email", claims.Email)
	c.Context().SetUserValue("user_id", claims.UserID)

	return c.Next()
}

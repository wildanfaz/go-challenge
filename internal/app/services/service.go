package services

import "github.com/gofiber/fiber/v2"

type Service func(c *fiber.Ctx) error

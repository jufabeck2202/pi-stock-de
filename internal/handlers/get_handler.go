package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type GetHandler struct {
	websiteService ports.WebsiteService
}

func NewGetHandler(websiteService ports.WebsiteService) *GetHandler {
	return &GetHandler{
		websiteService: websiteService,
	}
}

func (hdl *GetHandler) Get(c *fiber.Ctx) error {
	return c.JSON(hdl.websiteService.GetList())
}

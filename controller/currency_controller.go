package controller

import (
	"boletia/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CurrencyController struct {
	CurrencyService *services.CurrencyService
}

func NewCurrencyController(s *services.CurrencyService) *CurrencyController {
	return &CurrencyController{
		CurrencyService: s,
	}
}

func (c *CurrencyController) GetGurrencies(ctx echo.Context) error {
	currencies := c.CurrencyService.CurrencyData.Currencies

	return ctx.JSON(http.StatusOK, currencies)
}

package httpHandlers

import (
	"github.com/andrewchababi/pricecare/backend/calculator"
	"github.com/labstack/echo/v4"
)

type CalculatePanelRequest struct {
	TestIds []string `json:"testIds" binding:"required"`
}

func CalculatePanelPrice(c echo.Context) error {
	var req CalculatePanelRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, echo.Map{"error": "invalid JSON array"})
	}

	if len(req.TestIds) == 0 {
		return c.JSON(400, echo.Map{
			"error": "no test ids provided",
		})
	}

	price := calculator.CalculatePanelPrice(req.TestIds)

	return c.JSON(200, echo.Map{
		"total_price": price,
	})
}

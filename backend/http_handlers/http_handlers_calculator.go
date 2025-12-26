package httpHandlers

import (
	"github.com/andrewchababi/pricecare/backend/calculator"
	"github.com/labstack/echo/v4"
)

func CalculatePanelPrice(c echo.Context) error {
	var ids []string
	if err := c.Bind(&ids); err != nil {
		return c.JSON(400, echo.Map{"error": "invalid JSON array"})
	}

	if len(ids) == 0 {
		return c.JSON(400, echo.Map{
			"error": "no test ids provided",
		})
	}

	price := calculator.CalculatePanelPrice(ids)

	return c.JSON(200, echo.Map{
		"total_price": price,
	})
}

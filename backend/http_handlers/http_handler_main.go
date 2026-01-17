package httpHandlers

import (
	"github.com/a-h/templ"
	"github.com/andrewchababi/pricecare/backend/models"
	"github.com/andrewchababi/pricecare/web/pages"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, page templ.Component) error {
	return page.Render(c.Request().Context(), c.Response())
}

func LoginPage(c echo.Context) error {
	return render(c, pages.LoginPage())
}

func CalculatorPage(c echo.Context) error {
	user, _ := c.Get("user").(models.User)

	return render(c, pages.CalculatorPage(user))
}

func SettingsPage(c echo.Context) error {
	user, _ := c.Get("user").(models.User)
	return render(c, pages.SettingsPage(user))
}

// func renderAdminPage(c echo.Context, user models.User) error {
// 	data := driverPageData{
// 		Username:  user.Username,
// 		Locations: database.GetLocations(),
// 		Messages: chat.GetMessages(),
// 	}

// 	return c.Render(http.StatusOK, "admin.html", data)
// }

// func renderDriverPage(c echo.Context, user models.User) error {
// 	data := driverPageData{
// 		Username:  user.Username,
// 		Locations: database.GetLocations(),
// 		Messages: chat.GetMessages(),
// 	}

//     return c.Render(http.StatusOK, "driver.html", data)
// }

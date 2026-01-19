package pages

import (
	"fmt"

	"github.com/andrewchababi/pricecare/backend/models"
)

var pagePrefetches = map[models.Page][]string{
	models.PageLogin: {
		"/images/calculator/calculator_filled.png",
		"/images/sidebar/settings_filled.svg",
	},
	models.PageCalculator: {"/images/calculator/calculator_filled.png"},
	models.PageSettings:   {},
}

var pageStyles = map[models.Page][]string{
	models.PageLogin:      {"/styles/login.css"},
	models.PageCalculator: {"/styles/calculator.css"},
	models.PageSettings:   {"/styles/settings.css"},
}

var pagePreScripts = map[models.Page][]string{
	models.PageLogin:      {},
	models.PageCalculator: {},
	models.PageSettings:   {},
	// models.PageHistory:   {"/scripts/websocket.js", "https://unpkg.com/htmx.org@1.9.6"},
}

var pagePostScripts = map[models.Page][]string{
	models.PageLogin:      {"/scripts/login.js"},
	models.PageCalculator: {"/scripts/calculator.js"},
	models.PageSettings:   {"/scripts/settings.js"},
}

var pagePrefetchs = map[models.Page][]string{
	models.PageLogin:      {},
	models.PageCalculator: {},
	models.PageSettings:   {},
}

func conditionalClass(condition bool, result string) string {
	if condition {
		return fmt.Sprintf(" %s", result)
	} else {
		return ""
	}
}

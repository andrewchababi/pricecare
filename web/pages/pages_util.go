package pages

import (
	"fmt"

	"github.com/andrewchababi/pricecare/backend/models"
)

var pageStyles = map[models.Page][]string{
	models.PageLogin:      {"/styles/login.css"},
	models.PageCalculator: {"/styles/login.css"},
}

var pagePreScripts = map[models.Page][]string{
	models.PageLogin:      {},
	models.PageCalculator: {},
	// models.PageHistory:   {"/scripts/websocket.js", "https://unpkg.com/htmx.org@1.9.6"},
}

var pagePostScripts = map[models.Page][]string{
	models.PageLogin:      {"/scripts/login.js"},
	models.PageCalculator: {"/scripts/Calculator.js"},
}

var pagePrefetchs = map[models.Page][]string{
	models.PageLogin:      {},
	models.PageCalculator: {},
}

func conditionalClass(condition bool, result string) string {
	if condition {
		return fmt.Sprintf(" %s", result)
	} else {
		return ""
	}
}

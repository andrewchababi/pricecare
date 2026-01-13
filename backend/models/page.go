package models

type Page int

const (
	PageLogin Page = iota
	PageCalculator
)

var PageName = map[Page]string{
	PageLogin:      "Login",
	PageCalculator: "Calculator",
}

func GetPageName(page Page) string {
	return PageName[page]
}

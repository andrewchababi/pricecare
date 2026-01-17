package models

type Page int

const (
	PageLogin Page = iota
	PageCalculator
	PageSettings
)

var PageName = map[Page]string{
	PageLogin:      "Login",
	PageCalculator: "Calculator",
	PageSettings:   "Settings",
}

var pagePath = map[Page]string{
	PageLogin:      "/login",
	PageCalculator: "/calculator",
	PageSettings:   "/settings",
}

func GetPageName(page Page) string {
	return PageName[page]
}

func GetPagePath(page Page) string {
	return pagePath[page]
}

func GetVisiblePages(userType UserType) []Page {
	switch userType {
	case UserTypeStaff:
		return []Page{PageCalculator, PageSettings}
	case UserTypeAdminLab:
		return []Page{PageCalculator, PageSettings}
	default:
		return []Page{}
	}
}

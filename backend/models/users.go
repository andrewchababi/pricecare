package models

type UserType int

const (
	UserTypeNone UserType = iota
	UserTypeStaff
	UserTypeAdminLab
)

func GetUserTypeString(userType UserType) string {
	switch userType {
	case UserTypeStaff:
		return "staff"
	case UserTypeAdminLab:
		return "adminLab"
	default:
		return "none"
	}
}

func GetUserTypePrettyString(userType UserType) string {
	switch userType {
	case UserTypeStaff:
		return "Staff"
	case UserTypeAdminLab:
		return "AdminLab"
	default:
		return "none"
	}
}

type User struct {
	Username       string   `json:"username"`
	HashedPassword string   `json:"hashedPassword"`
	UserType       UserType `json:"userType"`
}

func NullUser() User {
	return User{
		Username:       "",
		HashedPassword: "",
		UserType:       UserTypeNone,
	}
}

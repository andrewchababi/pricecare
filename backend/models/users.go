package models

import "go.mongodb.org/mongo-driver/v2/bson"

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
	Id             bson.ObjectID `bson:"_id,omitempty"  json:"id"`
	Username       string        `bson:"username"      json:"username"`
	HashedPassword string        `bson:"hashedPassword" json:"hashedPassword"`
	UserType       UserType      `bson:"userType"      json:"userType"`
}

func NullUser() User {
	return User{
		Username:       "",
		HashedPassword: "",
		UserType:       UserTypeNone,
	}
}

type PublicUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	UserType string `json:"userType"`
}

func GetPublicUser(user User) PublicUser {
	return PublicUser{
		Id:       user.Id.Hex(),
		Username: user.Username,
		UserType: GetUserTypeString(user.UserType),
	}
}

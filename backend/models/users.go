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
	Username       string        `bson:"username"       json:"username"`
	HashedPassword string        `bson:"hashedPassword" json:"hashedPassword"`
	FirstName      string        `bson:"firstName"      json:"firstName"`
	LastName       string        `bson:"lastName"       json:"lastName"`
	UserType       UserType      `bson:"userType"       json:"userType"`
}

func NullUser() User {
	return User{
		Id:             bson.NilObjectID,
		Username:       "",
		HashedPassword: "",
		FirstName:      "",
		LastName:       "",
		UserType:       UserTypeNone,
	}
}

type PublicUser struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserType  string `json:"userType"`
}

func GetPublicUser(user User) PublicUser {
	return PublicUser{
		Id:        user.Id.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserType:  GetUserTypeString(user.UserType),
	}
}

func GetUserMap(userList []User) map[bson.ObjectID]User {
	userMap := make(map[bson.ObjectID]User, len(userList))
	for _, user := range userList {
		userMap[user.Id] = user
	}
	return userMap
}

func GetUserIds(users []User) []bson.ObjectID {
	ids := make([]bson.ObjectID, len(users))
	for i, user := range users {
		ids[i] = user.Id
	}
	return ids
}

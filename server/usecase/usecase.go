package usecase

import (
	"crud-grpc/server/model"
)

type UserClass struct{}

func NewUserClass() *UserClass {
	return &UserClass{}
}

var userData []model.User

func (uc *UserClass) AddUser(user model.User) []model.User {
	userData = append(userData, user)
	return userData
}

func (uc *UserClass) FindUserByID(id int) *model.User {
	if id >= len(userData) {
		return &model.User{}
	}
	return &userData[id]
}

func (uc *UserClass) GetAll() []model.User {
	return userData
}

// func (uc *UserClass) UpdateUser(id int, user model.User) response.Response {
// 	return nil
// }

// func (uc *UserClass) DeleteUser(id int) []model.User {

// 	if id >= len(userData) {
// 		return []model.User{}
// 	}

// 	userData = append(userData[:id], userData...)

// 	return userData
// }

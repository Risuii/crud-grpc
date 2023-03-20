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

func (uc *UserClass) UpdateUser(id int, user model.User) model.User {
	if id >= len(userData) {
		return model.User{}
	}

	userData[id] = user

	return user
}

func (uc *UserClass) DeleteUser(id int) {

	if id >= len(userData) {
	}

	userData = append(userData[:id], userData[id+1:]...)

}

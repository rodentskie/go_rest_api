package userService

import userModel "go-rest-api/src/models"

type UserService interface {
	CreateUser(*userModel.User) error
	GetSingleUser(*string) (*userModel.User, error)
	GetAllUser() ([]*userModel.User, error)
	UpdateUser(*userModel.User) error
	DeleteUser(*string) error
}

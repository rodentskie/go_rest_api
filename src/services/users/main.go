package userService

import (
	userModel "go-rest-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	CreateUser(*userModel.User) (primitive.ObjectID, error)
	GetSingleUser(*string) (*userModel.User, error)
	GetAllUser() ([]*userModel.User, error)
	UpdateUser(*string, *userModel.User) error
	DeleteUser(*string) error
}

package userService

import (
	"context"
	"errors"
	userModel "go-rest-api/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userDb *mongo.Collection
	ctx    context.Context
}

func InitUserService(userDb *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userDb: userDb,
		ctx:    ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *userModel.User) error {
	filterName := bson.D{{Key: "name", Value: user.Name}}
	filterEmail := bson.D{{Key: "email", Value: user.Email}}

	if err := u.userDb.FindOne(u.ctx, filterName).Decode(&user); err != mongo.ErrNoDocuments {
		return errors.New("User exist.")
	}

	if err := u.userDb.FindOne(u.ctx, filterEmail).Decode(&user); err != mongo.ErrNoDocuments {
		return errors.New("Email exist.")
	}

	_, err := u.userDb.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetSingleUser(id *string) (*userModel.User, error) {
	var user *userModel.User
	objectId, _ := primitive.ObjectIDFromHex(*id)
	query := bson.M{"_id": objectId}

	err := u.userDb.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAllUser() ([]*userModel.User, error) {
	var users []*userModel.User
	cur, err := u.userDb.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	defer cur.Close(u.ctx)

	for cur.Next(u.ctx) {
		var user userModel.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *userModel.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}

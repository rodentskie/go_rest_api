package userService

import (
	"context"
	"errors"
	functions "go-rest-api/src/functions"
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

func (u *UserServiceImpl) CreateUser(user *userModel.User) (primitive.ObjectID, string, error) {
	filterName := bson.D{{Key: "name", Value: user.Name}}
	filterEmail := bson.D{{Key: "email", Value: user.Email}}

	if err := u.userDb.FindOne(u.ctx, filterName).Decode(&user); err != mongo.ErrNoDocuments {
		return primitive.NilObjectID, "", errors.New("User exist.")
	}

	if err := u.userDb.FindOne(u.ctx, filterEmail).Decode(&user); err != mongo.ErrNoDocuments {
		return primitive.NilObjectID, "", errors.New("Email exist.")
	}

	res, err := u.userDb.InsertOne(u.ctx, user)
	oid, _ := res.InsertedID.(primitive.ObjectID)

	ss, err := functions.GenerateToken(oid.Hex())

	return oid, ss, err
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

func (u *UserServiceImpl) UpdateUser(id *string, user *userModel.User) error {
	objectId, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.M{"_id": objectId}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: user.Name}, primitive.E{Key: "email", Value: user.Email}, primitive.E{Key: "age", Value: user.Age}, primitive.E{Key: "hobby", Value: user.Hobby}}}}

	result, _ := u.userDb.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("User doesn't exist.")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(id *string) error {
	objectId, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.M{"_id": objectId}
	result, _ := u.userDb.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("User doesn't exist.")
	}
	return nil
}

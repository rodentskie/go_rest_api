package userModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hobby struct {
	Name        string `json:"name" bson:"name" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required"`
}

type User struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name" validate:"required"`
	Email string             `json:"email" bson:"email" validate:"required,email"`
	Age   int                `json:"age" bson:"age" validate:"required,gte=0,lte=130"`
	Hobby []*Hobby           `json:"hobby" bson:"hobby" validate:"required"`
}

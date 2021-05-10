package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Full_name  *string            `json:"full_name" validate:"required,min=5,max=100"`
	Email      *string            `json:"email" validate:"email,required"`
	Password   *string            `json:"Password" validate:"required,min=6"`
	Standard   *int               `json:"standard" validate:"required"`
	Session    *int               `json:"session" validate:"required"`
	Token      *string            `json:"token"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	User_id    string             `json:"user_id"`
	Admin      bool               `json:"admin"`
}

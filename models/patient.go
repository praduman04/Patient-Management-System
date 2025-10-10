package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" validate:"required,min=2"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Phone     string             `bson:"phone" json:"phone" validate:"required"`
	Age       int                `bson:"age" json:"age" validate:"required,min=0"`
	Gender    string             `bson:"gender" json:"gender" validate:"required,oneof=male female other"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

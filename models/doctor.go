package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doctor struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name" validate:"required,min=2"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	Specialty string             `bson:"speciality" json:"speciality" validate:"required"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

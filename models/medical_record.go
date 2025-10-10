package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MedicalRecord struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID primitive.ObjectID `bson:"patient_id" json:"patient_id" validate:"required"`
	Diagnosis string             `bson:"diagnosis" json:"diagnosis" validate:"required"`
	Treatment string             `bson:"treatment" json:"treatment"`
	Notes     string             `bson:"notes" json:"notes"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID primitive.ObjectID `bson:"patient_id" json:"patient_id" validate:"required"`
	DoctorID  primitive.ObjectID `bson:"doctor_id" json:"doctor_id" validate:"required"`
	Date      time.Time          `bson:"date" json:"date" validate:"required"`
	Status    string             `bson:"status" json:"status" validate:"required,oneof=scheduled completed canceled"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

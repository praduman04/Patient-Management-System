package repo

import (
	"context"
	"pms/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DoctorRepo struct {
	collection *mongo.Collection
}

func NewDoctorRepo(db *mongo.Database) *DoctorRepo {
	return &DoctorRepo{collection: db.Collection("doctor")}

}
func (r *DoctorRepo) CreateDoctor(ctx context.Context, doctor models.Doctor) (*models.Doctor, error) {
	id := primitive.NewObjectID()
	doctor.ID = id
	doctor.CreatedAt = time.Now()
	doctor.UpdatedAt = time.Now()

	if _, err := r.collection.InsertOne(ctx, doctor); err != nil {
		return nil, err

	}
	return &doctor, nil
}
func (r *DoctorRepo) GetById(ctx context.Context, id string) (*models.Doctor, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var doctor models.Doctor
	if err := r.collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&doctor); err != nil {
		return nil, err
	}
	return &doctor, nil

}
func (r *DoctorRepo) GetByEmail(ctx context.Context, email string) (*models.Doctor, error) {
	var doctor models.Doctor
	if err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&doctor); err != nil {
		return nil, err
	}
	return &doctor, nil
}
func (r *DoctorRepo) GetBySpecialty(ctx context.Context, specialty string) ([]models.Doctor, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"specialty": specialty})
	if err != nil {
		return nil, err
	}
	var doctors []models.Doctor
	for cursor.Next(ctx) {
		var doctor models.Doctor
		if err := cursor.Decode(&doctor); err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}
func (r *DoctorRepo) DeleteById(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if _, err := r.collection.DeleteOne(ctx, bson.M{"_id": objId}); err != nil {
		return err
	}
	return nil

}
func (r *DoctorRepo) GetAll(ctx context.Context) ([]models.Doctor, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var doctors []models.Doctor
	for cursor.Next(ctx) {
		var doctor models.Doctor
		if err := cursor.Decode(&doctor); err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)

	}
	return doctors, nil

}
func (r *DoctorRepo) UpdateDoctor(ctx context.Context, id string, updatedData bson.M) (*models.Doctor, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": updatedData,
	}
	_, err = r.collection.UpdateByID(ctx, objId, update)
	if err != nil {
		return nil, err
	}
	var doctor models.Doctor
	if err := r.collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&doctor); err != nil {
		return nil, err
	}
	return &doctor, nil

}

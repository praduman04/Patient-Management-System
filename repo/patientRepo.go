package repo

import (
	"context"
	"pms/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PatientRepo struct {
	collection *mongo.Collection
}

func NewPatientRepo(db *mongo.Database) *PatientRepo {
	return &PatientRepo{collection: db.Collection("patients")}
}
func (r *PatientRepo) Create(ctx context.Context, patient models.Patient) (*mongo.InsertOneResult, error) {
	patient.Id = primitive.NewObjectID()
	patient.CreatedAt = time.Now()
	patient.UpdatedAt = time.Now()
	return r.collection.InsertOne(ctx, patient)
}
func (r *PatientRepo) GetAll(ctx context.Context) ([]models.Patient, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var patients []models.Patient
	for cursor.Next(ctx) {
		var patient models.Patient
		if err := cursor.Decode(&patient); err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	return patients, nil
}
func (r *PatientRepo) GetById(ctx context.Context, objId string) (*models.Patient, error) {
	id, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return nil, err
	}
	var patient models.Patient
	err = r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&patient)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}
func (r *PatientRepo) DeleteById(ctx context.Context, objId string) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return nil, err
	}
	return r.collection.DeleteOne(ctx, bson.M{"_id": id})

}
func (r *PatientRepo) Update(ctx context.Context, id string, updatedData models.Patient) (*mongo.UpdateResult, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	updatedData.UpdatedAt = time.Now()
	update := bson.M{
		"$set": bson.M{
			"name":   updatedData.Name,
			"email":  updatedData.Email,
			"phone":  updatedData.Phone,
			"age":    updatedData.Age,
			"gender": updatedData.Gender,
		},
	}
	return r.collection.UpdateByID(ctx, bson.M{"_id": objId}, update)
}
func (r *PatientRepo) ExistByEmail(ctx context.Context, email string) (bool, error) {
	filter := bson.M{"email": email}
	opts := options.Count().SetLimit(1)
	count, err := r.collection.CountDocuments(ctx, filter, opts)
	if err != nil {
		return false, err
	}
	return count > 0, err

}

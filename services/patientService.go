package services

import (
	"context"
	"errors"
	"pms/models"
	"pms/repo"
	"pms/transformer"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type PatientService struct {
	repo *repo.PatientRepo
}

func NewPatientService(repo *repo.PatientRepo) *PatientService {
	return &PatientService{repo: repo}
}
func (s *PatientService) CreatePatient(ctx context.Context, patient models.Patient) error {
	exists, err := s.repo.ExistByEmail(ctx, patient.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("Email already exists")
	}
	_, err = s.repo.Create(ctx, patient)
	return err

}
func (s *PatientService) GetAll(ctx context.Context) ([]models.Patient, error) {
	return s.repo.GetAll(ctx)
}
func (s *PatientService) GetById(ctx context.Context, id string) (*models.Patient, error) {
	return s.repo.GetById(ctx, id)
}
func (s *PatientService) Update(ctx context.Context, id string, patient transformer.UpdatePatient) (*models.Patient, error) {
	update := bson.M{}
	if patient.Name != nil {
		update["name"] = patient.Name
	}
	if patient.Email != nil {
		update["email"] = patient.Email
	}
	if patient.Age != nil {
		update["age"] = patient.Age
	}
	if patient.Gender != nil {
		update["gender"] = patient.Gender
	}
	if patient.Phone != nil {
		update["phone"] = patient.Phone
	}
	update["updated_at"] = time.Now()
	if len(update) == 1 {
		return nil, errors.New("No fileds to update.")
	}

	return s.repo.Update(ctx, id, update)

}
func (s *PatientService) Delete(ctx context.Context, id string) error {
	_, err := s.repo.DeleteById(ctx, id)
	return err
}

package service

import (
	"context"
	"fmt"

	"github.com/jcromanu/demo/internal/errors"
	"github.com/jcromanu/demo/pkg/entities"
	"github.com/jcromanu/demo/pkg/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repository: repo,
	}
}

func (ser *Service) CreateCar(ctx context.Context, car entities.Car) (string, *errors.HttpError) {
	name := fmt.Sprintf("Car %s created", car.Name)
	fmt.Print(name)
	return name, nil
}

func (ser *Service) DeleteCar(ctx context.Context, car entities.Car) *errors.HttpError {
	return nil
}

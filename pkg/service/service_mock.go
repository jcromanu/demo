package service

import (
	"context"

	"github.com/jcromanu/demo/internal/errors"
	"github.com/jcromanu/demo/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) CreateCar(ctx context.Context, car entities.Car) (string, *errors.HttpError) {
	args := s.Called(ctx, car)
	name := args.Get(0).(string)
	error := args.Get(1).(errors.HttpError)
	return name, &error
}

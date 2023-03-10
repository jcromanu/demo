package service

import (
	"context"
	"testing"

	"github.com/jcromanu/demo/pkg/entities"
	"github.com/jcromanu/demo/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateCar(t *testing.T) {
	type args struct {
		car *entities.Car
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Fine error ",
			args: args{
				car: &entities.Car{
					Name: "first_name",
				},
			},
			want:    "Car first_name created",
			wantErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := NewService(*repository.NewRepository(nil))
			carName, err := service.CreateCar(context.Background(), *test.args.car)
			assert.Equal(t, carName, test.want)
			if err != nil {
				assert.Equal(t, err, test.wantErr)
			}
		})
	}
}

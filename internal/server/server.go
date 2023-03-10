package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcromanu/demo/internal/errors"
	"github.com/jcromanu/demo/pkg/entities"
)

type CarService interface {
	CreateCar(context.Context, entities.Car) (string, *errors.HttpError)
	DeleteCar(context.Context, entities.Car) *errors.HttpError
}

type Server struct {
	router  *mux.Router
	service CarService
}

func NewServer(router *mux.Router, service CarService) *Server {
	return &Server{
		router:  router,
		service: service,
	}
}

func (r *Server) SetRoutes() {
	r.router.HandleFunc("/create", r.create()).Methods(http.MethodPost)
}

func (r *Server) create() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var car entities.Car
		json.NewDecoder(req.Body).Decode(&car)
		message, err := r.service.CreateCar(req.Context(), car)
		if err != nil {
			rw.WriteHeader(err.StatusCode)
			json.NewEncoder(rw).Encode(err.Message)
		} else {
			rw.WriteHeader(http.StatusAccepted)
			json.NewEncoder(rw).Encode(message)
		}
	}
}

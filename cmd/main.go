package main

import (
	"database/sql"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jcromanu/demo/internal/server"
	"github.com/jcromanu/demo/pkg/repository"
	"github.com/jcromanu/demo/pkg/service"
	"go.uber.org/zap"
)

func main() {
	//crear conexión a mysql
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	defer logger.Sync()
	mysqlConfig := mysql.Config{
		User:   "",
		Passwd: "",
		Net:    "",
		Addr:   "",
		DBName: "",
	}
	db, _ := sql.Open("mysql", mysqlConfig.FormatDSN())
	repo := repository.NewRepository(db)
	service := service.NewService(*repo)
	mux := mux.NewRouter()
	srv := server.NewServer(mux, service)
	srv.SetRoutes()
	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}
	errChan := make(chan error)
	sugar.Info("Starting server")
	go func() {
		errChan <- server.ListenAndServe()
	}()
	defer sugar.Info("Server stopped ")
	<-errChan
}

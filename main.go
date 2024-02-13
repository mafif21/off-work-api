package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-cuti-karyawan/app"
	"rest-api-cuti-karyawan/controller"
	"rest-api-cuti-karyawan/helper"
	"rest-api-cuti-karyawan/repository"
	"rest-api-cuti-karyawan/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
}

func main() {
	db := app.Database()
	validate := validator.New()

	newRepository := repository.NewOffWorkRepositoryImpl()
	newService := service.NewOffWorkServiceImpl(newRepository, db, validate)
	newController := controller.NewOffWorkControllerImpl(newService)

	router := app.NewRouter(newController)
	server := NewServer(router)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

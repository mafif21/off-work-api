package main

import (
	"net/http"
	"rest-api-cuti-karyawan/app"
	"rest-api-cuti-karyawan/controller"
	"rest-api-cuti-karyawan/exception"
	"rest-api-cuti-karyawan/helper"
	"rest-api-cuti-karyawan/repository"
	"rest-api-cuti-karyawan/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.Database()
	validate := validator.New()

	newRepository := repository.NewOffWorkRepositoryImpl()
	newService := service.NewOffWorkServiceImpl(newRepository, db, validate)
	newController := controller.NewOffWorkControllerImpl(newService)

	router := httprouter.New()

	router.GET("/api/offwork", newController.GetAll)
	router.GET("/api/offwork/:id", newController.FindDataById)
	router.POST("/api/offwork/create", newController.Create)
	router.PATCH("/api/offwork/changestatus/:id", newController.UpdateStatus)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

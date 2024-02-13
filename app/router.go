package app

import (
	"github.com/julienschmidt/httprouter"
	"rest-api-cuti-karyawan/controller"
	"rest-api-cuti-karyawan/exception"
)

func NewRouter(newController controller.OffWorkController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/offwork", newController.GetAll)
	router.GET("/api/offwork/:id", newController.FindDataById)
	router.POST("/api/offwork/create", newController.Create)
	router.PATCH("/api/offwork/changestatus/:id", newController.UpdateStatus)

	router.PanicHandler = exception.ErrorHandler
	return router
}

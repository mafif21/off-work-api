package controller

import (
	"net/http"
	"rest-api-cuti-karyawan/helper"
	"rest-api-cuti-karyawan/model/web"
	"rest-api-cuti-karyawan/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type OffWorkControllerImpl struct {
	Service service.OffWorkService
}

func NewOffWorkControllerImpl(service service.OffWorkService) OffWorkController {
	return &OffWorkControllerImpl{Service: service}
}

func (controller OffWorkControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	offWorkRequest := web.OffWorkRequest{}
	helper.ReadFromBodyRequest(request, &offWorkRequest)

	offWorkResponse := controller.Service.Create(request.Context(), offWorkRequest)
	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Message: "request off work has been created",
		Data:    offWorkResponse,
	}

	helper.WriteToBodyResponse(writer, webResponse)
}

func (controller OffWorkControllerImpl) FindDataById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	offWorkResponse := controller.Service.FindDataById(request.Context(), idInt)
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "success get data by id " + id,
		Data:    offWorkResponse,
	}

	helper.WriteToBodyResponse(writer, webResponse)
}

func (controller OffWorkControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	offWorkResponse := controller.Service.GetAll(request.Context())
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "success get all datas",
		Data:    offWorkResponse,
	}
	helper.WriteToBodyResponse(writer, webResponse)
}

func (controller OffWorkControllerImpl) UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	offWorkRequest := web.OffworkUpdateRequest{}
	helper.ReadFromBodyRequest(request, &offWorkRequest)

	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)
	offWorkRequest.Id = idInt

	offWorkResponse := controller.Service.UpdateStatus(request.Context(), offWorkRequest)
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "off work has been updated",
		Data:    offWorkResponse,
	}

	helper.WriteToBodyResponse(writer, webResponse)
}

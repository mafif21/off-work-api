package exception

import (
	"net/http"
	"rest-api-cuti-karyawan/helper"
	"rest-api-cuti-karyawan/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Message: "DATA NOT FOUND",
			Data:    exception.Error,
		}

		helper.WriteToBodyResponse(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:    http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
		Data:    err,
	}

	helper.WriteToBodyResponse(writer, webResponse)
}

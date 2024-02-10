package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OffWorkController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindDataById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

package helper

import (
	"encoding/json"
	"net/http"
	"rest-api-cuti-karyawan/model/web"
)

func ReadFromBodyRequest(request *http.Request, data interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(data)
	PanicIfError(err)
}

func WriteToBodyResponse(writer http.ResponseWriter, data web.WebResponse) {
	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(data)
	PanicIfError(err)
}

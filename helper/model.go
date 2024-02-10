package helper

import (
	"rest-api-cuti-karyawan/model/entity"
	"rest-api-cuti-karyawan/model/web"
)

func ToOffWorkResponse(offwork entity.OffWork) web.OffWorkResponse {
	return web.OffWorkResponse{
		Id:         offwork.Id,
		Name:       offwork.Name,
		Position:   offwork.Position,
		Start_date: offwork.Start_date,
		End_date:   offwork.End_date,
		Status:     offwork.Status,
	}
}

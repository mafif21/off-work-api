package service

import (
	"context"
	"rest-api-cuti-karyawan/model/web"
)

type OffWorkService interface {
	Create(ctx context.Context, request web.OffWorkRequest) web.OffWorkResponse
	FindDataById(ctx context.Context, id int) web.OffWorkResponse
	GetAll(ctx context.Context) []web.OffWorkResponse
	UpdateStatus(ctx context.Context, request web.OffworkUpdateRequest) web.OffWorkResponse
}

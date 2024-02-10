package service

import (
	"context"
	"database/sql"
	"rest-api-cuti-karyawan/exception"
	"rest-api-cuti-karyawan/helper"
	"rest-api-cuti-karyawan/model/entity"
	"rest-api-cuti-karyawan/model/web"
	"rest-api-cuti-karyawan/repository"

	"github.com/go-playground/validator/v10"
)

type OffWorkServiceImpl struct {
	Repository repository.OffWorkRepository
	Db         *sql.DB
	Validate   *validator.Validate
}

func NewOffWorkServiceImpl(repository repository.OffWorkRepository, db *sql.DB, validate *validator.Validate) OffWorkService {
	return &OffWorkServiceImpl{Repository: repository, Db: db, Validate: validate}
}

func (service *OffWorkServiceImpl) Create(ctx context.Context, request web.OffWorkRequest) web.OffWorkResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	offWork := entity.OffWork{
		Name:       request.Name,
		Position:   request.Position,
		Start_date: request.Start_date,
		End_date:   request.End_date,
		Status:     request.Status,
	}

	newOffWork := service.Repository.Save(ctx, tx, offWork)
	return helper.ToOffWorkResponse(newOffWork)

}

func (service *OffWorkServiceImpl) FindDataById(ctx context.Context, id int) web.OffWorkResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	offWork, err := service.Repository.FindDataById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToOffWorkResponse(offWork)
}

func (service *OffWorkServiceImpl) GetAll(ctx context.Context) []web.OffWorkResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	var offWorksResponse []web.OffWorkResponse
	offWorks := service.Repository.GetAll(ctx, tx)

	for _, offWork := range offWorks {
		offWorksResponse = append(offWorksResponse, helper.ToOffWorkResponse(offWork))
	}

	return offWorksResponse
}

func (service *OffWorkServiceImpl) UpdateStatus(ctx context.Context, request web.OffworkUpdateRequest) web.OffWorkResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	offWork, err := service.Repository.FindDataById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	offWork.Status = request.Status

	offWorkUpdated := service.Repository.UpdateStatus(ctx, tx, offWork)
	return helper.ToOffWorkResponse(offWorkUpdated)
}

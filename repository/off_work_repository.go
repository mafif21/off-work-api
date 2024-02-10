package repository

import (
	"context"
	"database/sql"
	"rest-api-cuti-karyawan/model/entity"
)

type OffWorkRepository interface {
	Save(ctx context.Context, tx *sql.Tx, offWork entity.OffWork) entity.OffWork
	FindDataById(ctx context.Context, tx *sql.Tx, id int) (entity.OffWork, error)
	GetAll(ctx context.Context, tx *sql.Tx) []entity.OffWork
	UpdateStatus(ctx context.Context, tx *sql.Tx, offWork entity.OffWork) entity.OffWork
}

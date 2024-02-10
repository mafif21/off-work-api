package repository

import (
	"context"
	"database/sql"
	"errors"
	"rest-api-cuti-karyawan/helper"
	"rest-api-cuti-karyawan/model/entity"
)

type OffWorkRepositoryImpl struct {
}

func NewOffWorkRepositoryImpl() *OffWorkRepositoryImpl {
	return &OffWorkRepositoryImpl{}
}

func (repository OffWorkRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, offWork entity.OffWork) entity.OffWork {
	query := "INSERT INTO off_works (name, position, start_date, end_date, status) VALUES (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, query, offWork.Name, offWork.Position, offWork.Start_date, offWork.End_date, offWork.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	offWork.Id = int(id)
	return offWork
}

func (repository OffWorkRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []entity.OffWork {
	query := "SELECT id, name, position, start_date, end_date, status FROM off_works"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var off_works []entity.OffWork
	for rows.Next() {
		offWork := entity.OffWork{}
		err := rows.Scan(&offWork.Id, &offWork.Name, &offWork.Position, &offWork.Start_date, &offWork.End_date, &offWork.Status)
		helper.PanicIfError(err)

		off_works = append(off_works, offWork)
	}

	return off_works
}

func (repository OffWorkRepositoryImpl) FindDataById(ctx context.Context, tx *sql.Tx, id int) (entity.OffWork, error) {
	query := "SELECT id, name, position, start_date, end_date, status FROM off_works WHERE id=?"
	rows, err := tx.QueryContext(ctx, query, id)
	helper.PanicIfError(err)
	defer rows.Close()

	offWork := entity.OffWork{}

	if rows.Next() {
		err := rows.Scan(&offWork.Id, &offWork.Name, &offWork.Position, &offWork.Start_date, &offWork.End_date, &offWork.Status)
		helper.PanicIfError(err)
		return offWork, nil
	} else {
		return offWork, errors.New("data not found")
	}
}

func (repository OffWorkRepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, offWork entity.OffWork) entity.OffWork {
	query := "UPDATE off_works set status = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, offWork.Status, offWork.Id)
	helper.PanicIfError(err)

	return offWork
}

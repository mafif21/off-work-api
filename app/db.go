package app

import (
	"database/sql"
	"rest-api-cuti-karyawan/helper"
	"time"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/offwork_restapi?parseTime=true")
	helper.PanicIfError(err)

	//migrate -database "mysql://root:@unix(/cloudsql/dietin-capstone:asia-southeast2:pratugas-akhir)/offwork" -path db/migrations up

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

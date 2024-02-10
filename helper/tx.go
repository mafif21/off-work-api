package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	errMessage := recover()
	if errMessage != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
		panic(errMessage)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}

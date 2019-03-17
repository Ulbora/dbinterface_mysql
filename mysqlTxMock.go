package mysql

import (
	"database/sql"
)

//MyDbTxMock MyDbTxMock
type MyDbTxMock struct {
	Tx *sql.Tx
	//MockSuccess bool
	//MockID      int64
	MyDBMock *MyDBMock
}

//Insert Insert
func (t *MyDbTxMock) Insert(query string, args ...interface{}) (bool, int64) {
	return t.MyDBMock.MockSuccess, t.MyDBMock.MockID
}

//Update Update
func (t *MyDbTxMock) Update(query string, args ...interface{}) bool {
	return t.MyDBMock.MockSuccess
}

//Delete Delete
func (t *MyDbTxMock) Delete(query string, args ...interface{}) bool {
	return t.MyDBMock.MockSuccess
}

//Commit Commit
func (t *MyDbTxMock) Commit() bool {
	return t.MyDBMock.MockSuccess
}

//Rollback Rollback
func (t *MyDbTxMock) Rollback() bool {
	return t.MyDBMock.MockSuccess
}

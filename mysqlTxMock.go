package mysql

import (
	"database/sql"
)

//MyDbTxMock MyDbTxMock
type MyDbTxMock struct {
	Tx       *sql.Tx
	MyDBMock *MyDBMock
}

//Insert Insert
func (t *MyDbTxMock) Insert(query string, args ...interface{}) (bool, int64) {
	var rtn = false
	var id int64
	if !t.MyDBMock.mockInsertSuccess1Used {
		t.MyDBMock.mockInsertSuccess1Used = true
		rtn = t.MyDBMock.MockInsertSuccess1
		id = t.MyDBMock.MockInsertID1
	} else if !t.MyDBMock.mockInsertSuccess2Used {
		t.MyDBMock.mockInsertSuccess2Used = true
		rtn = t.MyDBMock.MockInsertSuccess2
		id = t.MyDBMock.MockInsertID2
	} else if !t.MyDBMock.mockInsertSuccess3Used {
		t.MyDBMock.mockInsertSuccess3Used = true
		rtn = t.MyDBMock.MockInsertSuccess3
		id = t.MyDBMock.MockInsertID3
	} else if !t.MyDBMock.mockInsertSuccess4Used {
		t.MyDBMock.mockInsertSuccess4Used = true
		rtn = t.MyDBMock.MockInsertSuccess4
		id = t.MyDBMock.MockInsertID4
	} else if !t.MyDBMock.mockInsertSuccess5Used {
		t.MyDBMock.mockInsertSuccess5Used = true
		rtn = t.MyDBMock.MockInsertSuccess5
		id = t.MyDBMock.MockInsertID5
	} else if !t.MyDBMock.mockInsertSuccess6Used {
		t.MyDBMock.mockInsertSuccess6Used = true
		rtn = t.MyDBMock.MockInsertSuccess6
		id = t.MyDBMock.MockInsertID6
	} else if !t.MyDBMock.mockInsertSuccess7Used {
		t.MyDBMock.mockInsertSuccess7Used = true
		rtn = t.MyDBMock.MockInsertSuccess7
		id = t.MyDBMock.MockInsertID7
	} else if !t.MyDBMock.mockInsertSuccess8Used {
		t.MyDBMock.mockInsertSuccess8Used = true
		rtn = t.MyDBMock.MockInsertSuccess8
		id = t.MyDBMock.MockInsertID8
	}
	return rtn, id
}

//Update Update
func (t *MyDbTxMock) Update(query string, args ...interface{}) bool {
	var rtn = false
	if !t.MyDBMock.mockUpdateSuccess1Used {
		t.MyDBMock.mockUpdateSuccess1Used = true
		rtn = t.MyDBMock.MockUpdateSuccess1
	} else if !t.MyDBMock.mockUpdateSuccess2Used {
		t.MyDBMock.mockUpdateSuccess2Used = true
		rtn = t.MyDBMock.MockUpdateSuccess2
	} else if !t.MyDBMock.mockUpdateSuccess3Used {
		t.MyDBMock.mockUpdateSuccess3Used = true
		rtn = t.MyDBMock.MockUpdateSuccess3
	} else if !t.MyDBMock.mockUpdateSuccess4Used {
		t.MyDBMock.mockUpdateSuccess4Used = true
		rtn = t.MyDBMock.MockUpdateSuccess4
	}
	return rtn
}

//Delete Delete
func (t *MyDbTxMock) Delete(query string, args ...interface{}) bool {
	var rtn = false
	if !t.MyDBMock.mockDeleteSuccess1Used {
		t.MyDBMock.mockDeleteSuccess1Used = true
		rtn = t.MyDBMock.MockDeleteSuccess1
	} else if !t.MyDBMock.mockDeleteSuccess2Used {
		t.MyDBMock.mockDeleteSuccess2Used = true
		rtn = t.MyDBMock.MockDeleteSuccess2
	} else if !t.MyDBMock.mockDeleteSuccess3Used {
		t.MyDBMock.mockDeleteSuccess3Used = true
		rtn = t.MyDBMock.MockDeleteSuccess3
	} else if !t.MyDBMock.mockDeleteSuccess4Used {
		t.MyDBMock.mockDeleteSuccess4Used = true
		rtn = t.MyDBMock.MockDeleteSuccess4
	} else if !t.MyDBMock.mockDeleteSuccess5Used {
		t.MyDBMock.mockDeleteSuccess5Used = true
		rtn = t.MyDBMock.MockDeleteSuccess5
	} else if !t.MyDBMock.mockDeleteSuccess6Used {
		t.MyDBMock.mockDeleteSuccess6Used = true
		rtn = t.MyDBMock.MockDeleteSuccess6
	} else if !t.MyDBMock.mockDeleteSuccess7Used {
		t.MyDBMock.mockDeleteSuccess7Used = true
		rtn = t.MyDBMock.MockDeleteSuccess7
	} else if !t.MyDBMock.mockDeleteSuccess8Used {
		t.MyDBMock.mockDeleteSuccess8Used = true
		rtn = t.MyDBMock.MockDeleteSuccess8
	}
	return rtn
}

//Commit Commit
func (t *MyDbTxMock) Commit() bool {
	return t.MyDBMock.MockCommitSuccess
}

//Rollback Rollback
func (t *MyDbTxMock) Rollback() bool {
	return t.MyDBMock.MockRollbackSuccess
}

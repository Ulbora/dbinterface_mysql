package mysql

import (
	"database/sql"
	di "github.com/Ulbora/dbinterface"
	//"log"

	_ "github.com/go-sql-driver/mysql"
)

//MyDBMock MyDBMock
type MyDBMock struct {
	Host                   string
	User                   string
	Password               string
	Database               string
	db                     *sql.DB
	err                    error
	MockConnectSuccess     bool
	MockCloseSuccess       bool
	MockCommitSuccess      bool
	MockRollbackSuccess    bool
	MockInsertSuccess1     bool
	MockInsertSuccess2     bool
	MockInsertSuccess3     bool
	MockInsertSuccess4     bool
	mockInsertSuccess1Used bool
	mockInsertSuccess2Used bool
	mockInsertSuccess3Used bool
	mockInsertSuccess4Used bool
	MockInsertID1          int64
	MockInsertID2          int64
	MockInsertID3          int64
	MockInsertID4          int64
	MockUpdateSuccess1     bool
	MockUpdateSuccess2     bool
	MockUpdateSuccess3     bool
	MockUpdateSuccess4     bool
	mockUpdateSuccess1Used bool
	mockUpdateSuccess2Used bool
	mockUpdateSuccess3Used bool
	mockUpdateSuccess4Used bool
	MockDeleteSuccess1     bool
	MockDeleteSuccess2     bool
	MockDeleteSuccess3     bool
	MockDeleteSuccess4     bool
	mockDeleteSuccess1Used bool
	mockDeleteSuccess2Used bool
	mockDeleteSuccess3Used bool
	mockDeleteSuccess4Used bool
	MockTestRow            *di.DbRow
	MockRow                *di.DbRow
	MockRows               *di.DbRows
}

//Connect Connect
func (m *MyDBMock) Connect() bool {
	return m.MockConnectSuccess
}

//BeginTransaction BeginTransaction
func (m *MyDBMock) BeginTransaction() di.Transaction {
	var trans di.Transaction
	var mtx MyDbTxMock
	mtx.MyDBMock = m
	trans = &mtx
	return trans
}

//Test Test
func (m *MyDBMock) Test(query string, args ...interface{}) *di.DbRow {
	return m.MockTestRow
}

//Insert Insert
func (m *MyDBMock) Insert(query string, args ...interface{}) (bool, int64) {
	var rtn = false
	var id int64
	if !m.mockInsertSuccess1Used {
		m.mockInsertSuccess1Used = true
		rtn = m.MockInsertSuccess1
		id = m.MockInsertID1
	} else if !m.mockInsertSuccess2Used {
		m.mockInsertSuccess2Used = true
		rtn = m.MockInsertSuccess2
		id = m.MockInsertID2
	} else if !m.mockInsertSuccess3Used {
		m.mockInsertSuccess3Used = true
		rtn = m.MockInsertSuccess3
		id = m.MockInsertID3
	} else if !m.mockInsertSuccess4Used {
		m.mockInsertSuccess4Used = true
		rtn = m.MockInsertSuccess4
		id = m.MockInsertID4
	}
	return rtn, id
}

//Update Update
func (m *MyDBMock) Update(query string, args ...interface{}) bool {
	var rtn = false
	if !m.mockUpdateSuccess1Used {
		m.mockUpdateSuccess1Used = true
		rtn = m.MockUpdateSuccess1
	} else if !m.mockUpdateSuccess2Used {
		m.mockUpdateSuccess2Used = true
		rtn = m.MockUpdateSuccess2
	} else if !m.mockUpdateSuccess3Used {
		m.mockUpdateSuccess3Used = true
		rtn = m.MockUpdateSuccess3
	} else if !m.mockUpdateSuccess4Used {
		m.mockUpdateSuccess4Used = true
		rtn = m.MockUpdateSuccess4
	}
	return rtn
}

//Get Get
func (m *MyDBMock) Get(query string, args ...interface{}) *di.DbRow {
	return m.MockRow
}

//GetList GetList
func (m *MyDBMock) GetList(query string, args ...interface{}) *di.DbRows {
	return m.MockRows
}

//Delete Delete
func (m *MyDBMock) Delete(query string, args ...interface{}) bool {
	var rtn = false
	if !m.mockDeleteSuccess1Used {
		m.mockDeleteSuccess1Used = true
		rtn = m.MockDeleteSuccess1
	} else if !m.mockDeleteSuccess2Used {
		m.mockDeleteSuccess2Used = true
		rtn = m.MockDeleteSuccess2
	} else if !m.mockDeleteSuccess3Used {
		m.mockDeleteSuccess3Used = true
		rtn = m.MockDeleteSuccess3
	} else if !m.mockDeleteSuccess4Used {
		m.mockDeleteSuccess4Used = true
		rtn = m.MockDeleteSuccess4
	}
	return rtn
}

//Close Close
func (m *MyDBMock) Close() bool {
	return m.MockCloseSuccess
}

package mysql

import (
	//"database/sql"
	"fmt"
	"testing"

	di "github.com/Ulbora/dbinterface"
)

var tx1m di.Transaction
var mdbtm MyDBMock
var dbitm di.Database

var idtx1m int64

func TestMyDbTxMock_Connect(t *testing.T) {

	mdbtm.Host = "localhost:3306"
	mdbtm.User = "admin"
	mdbtm.Password = "admin"
	mdbtm.Database = "testdb"
	mdbtm.MockSuccess = true
	dbitm = &mdbtm
	suc := dbitm.Connect()
	if !suc {
		fmt.Println("Could not connect: ")
	}

}

func TestMyDbTxMock_Insert(t *testing.T) {
	tx1m = dbitm.BeginTransaction()
	mdbtm.MockID = 1

	var q = "insert into test (name, address) values(?, ?)"
	var a []interface{}
	a = append(a, "test insert 1", "123 main st")

	suc, id := tx1m.Insert(q, a...)
	if !suc || id < 1 {
		t.Fail()
	} else {
		idtx1m = id
		fmt.Println("inserted id: ", id)
	}
	tx1m.Commit()
}

func TestMyDbTxMock_Insertfail(t *testing.T) {
	tx1m = dbitm.BeginTransaction()
	mdbtm.MockSuccess = false
	var q = "insert into test (name, address) values(?, ?)"
	var a []interface{}
	//a = append(a, "test insert 1", "123 main st")

	suc, id := tx1m.Insert(q, a...)
	if suc {
		t.Fail()
	} else {
		fmt.Println("inserted id: ", id)
	}
	tx1m.Commit()
}

func TestMyDbTxMock_Update(t *testing.T) {
	tx1xx := dbitm.BeginTransaction()
	mdbtm.MockSuccess = true
	var q = "update test set name = ? , address = ? where id = ? "
	var a []interface{}
	a = append(a, "test insert 2", "123456 main st", idtx1)
	suc := tx1xx.Update(q, a...)
	if !suc {
		t.Fail()
	}
	tx1xx.Commit()
}

func TestMyDbTxMock_Updatefail(t *testing.T) {
	tx1xx := dbitm.BeginTransaction()
	mdbtm.MockSuccess = false
	var q = "update test11 set name = ? , address = ? where id = ? "
	var a []interface{}
	a = append(a, "test insert 2", "123456 main st", idtx1)
	suc := tx1xx.Update(q, a...)
	if suc {
		t.Fail()
	}
	tx1xx.Commit()
}

func TestMyDbTxMock_Delete(t *testing.T) {
	tx1xx := dbitm.BeginTransaction()
	mdbtm.MockSuccess = true
	var q = "delete from test where id = ? "
	var a []interface{}
	a = append(a, idtx1)
	suc := tx1xx.Delete(q, a...)
	if !suc {
		t.Fail()
	}
	tx1xx.Commit()
}

func TestMyDbTxMock_Rollback(t *testing.T) {
	tx1xx := dbitm.BeginTransaction()
	suc := tx1xx.Rollback()
	if !suc {
		t.Fail()
	}
	tx1xx.Commit()
}

func TestMyDbTxMock_Deletefail(t *testing.T) {
	tx1xx := dbitm.BeginTransaction()
	mdbtm.MockSuccess = false
	var q = "delete from test22 where id = ? "
	var a []interface{}
	a = append(a, idtx1)
	suc := tx1xx.Delete(q, a...)
	if suc {
		t.Fail()
	}
	tx1xx.Commit()
}

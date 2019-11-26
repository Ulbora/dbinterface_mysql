package mysql

import (
	//"database/sql"
	"fmt"
	"testing"

	di "github.com/Ulbora/dbinterface"
)

var tx1 di.Transaction
var mdbt MyDB
var dbit di.Database

var idtx1 int64

func TestMyDbTx_Connect(t *testing.T) {

	mdbt.Host = "localhost:3306"
	mdbt.User = "admin"
	mdbt.Password = "admin"
	mdbt.Database = "testdb"
	//dbit = &mdbt
	dbit = mdbt.GetNewDatabase()
	suc := dbit.Connect()
	if !suc {
		fmt.Println("Could not connect: ")
	}

}

func TestMyDbTx_Insert(t *testing.T) {
	tx1 = dbit.BeginTransaction()

	var q = "insert into test (name, address) values(?, ?)"
	var a []interface{}
	a = append(a, "test insert 1", "123 main st")

	suc, id := tx1.Insert(q, a...)
	if !suc || id < 1 {
		t.Fail()
	} else {
		idtx1 = id
		fmt.Println("inserted id: ", id)
	}
	tx1.Commit()
}

func TestMyDbTx_Insertfail(t *testing.T) {
	tx1 = dbit.BeginTransaction()

	var q = "insert into test (name, address) values(?, ?)"
	var a []interface{}
	//a = append(a, "test insert 1", "123 main st")

	suc, id := tx1.Insert(q, a...)
	if suc {
		t.Fail()
	} else {
		fmt.Println("inserted id: ", id)
	}
	tx1.Commit()
}

func TestMyDbTx_Update(t *testing.T) {
	tx1xx := dbit.BeginTransaction()
	var q = "update test set name = ? , address = ? where id = ? "
	var a []interface{}
	a = append(a, "test insert 2", "123456 main st", idtx1)
	suc := tx1xx.Update(q, a...)
	if !suc {
		t.Fail()
	}
	tx1xx.Commit()
}

func TestMyDbTx_Updatefail(t *testing.T) {
	tx1xx := dbit.BeginTransaction()
	var q = "update test11 set name = ? , address = ? where id = ? "
	var a []interface{}
	a = append(a, "test insert 2", "123456 main st", idtx1)
	suc := tx1xx.Update(q, a...)
	if suc {
		t.Fail()
	}
	tx1xx.Commit()
}

func TestMyDbTx_Updatefail2(t *testing.T) {
	tx1xx := dbit.BeginTransaction()
	var q = "update test set name = ? , address = ? where id = ? "
	var a []interface{}
	//a = append(a, "test insert 2", "123456 main st", idtx1)
	suc := tx1xx.Update(q, a...)
	if suc {
		t.Fail()
	}
	tx1xx.Commit()
}

func TestMyDbTx_Delete(t *testing.T) {
	tx1xx := dbit.BeginTransaction()
	var q = "delete from test where id = ? "
	var a []interface{}
	a = append(a, idtx1)
	suc := tx1xx.Delete(q, a...)
	if !suc {
		t.Fail()
	}
	tx1xx.Commit()
}

func TestMyDbTx_Deletefail(t *testing.T) {
	tx1xx := dbit.BeginTransaction()
	var q = "delete from test22 where id = ? "
	var a []interface{}
	a = append(a, idtx1)
	suc := tx1xx.Delete(q, a...)
	if suc {
		t.Fail()
	}
	tx1xx.Commit()
}

func TestMyDbTx_Deletefail2(t *testing.T) {
	tx1xx := dbit.BeginTransaction()
	var q = "delete from test where id = ? "
	var a []interface{}
	//a = append(a, idtx1)
	suc := tx1xx.Delete(q, a...)
	if suc {
		t.Fail()
	}
	tx1xx.Commit()
}

func TestMyDbTx_Insert2(t *testing.T) {
	tx1 = dbit.BeginTransaction()

	var q = "insert into test (name, address) values(?, ?)"
	var a []interface{}
	a = append(a, "test insert 111", "123 main st")

	suc, id := tx1.Insert(q, a...)
	if !suc || id < 1 {
		t.Fail()
	} else {
		idtx1 = id
		fmt.Println("inserted id: ", id)
	}
	// tx1.Commit()
}

func TestMyDbTx_Insert3(t *testing.T) {
	//tx1 = mdbt.BeginTransaction()

	var q = "insert into test22 (name, address) values(?, ?)"
	var a []interface{}
	a = append(a, "test insert 111", "123 main st")

	suc, id := tx1.Insert(q, a...)
	if !suc || id < 1 {
		rbsuc := tx1.Rollback()
		if !rbsuc {
			t.Fail()
		}
	} else {
		t.Fail()
	}
	// tx1.Commit()
}
func TestMyDbTx_Close(t *testing.T) {
	suc := dbit.Close()
	if !suc {
		t.Fail()
	}
}

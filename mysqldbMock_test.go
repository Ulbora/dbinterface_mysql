package mysql

import (
	"fmt"
	"strconv"
	//"database/sql"
	di "github.com/Ulbora/dbinterface"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var dbim di.Database

func TestMyDBMock_Connect(t *testing.T) {
	var mdb MyDBMock
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin1"
	mdb.Database = "testdb1"
	mdb.MockConnectSuccess = true
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	suc := dbim.Connect()
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Test(t *testing.T) {
	var mdb MyDBMock
	var rtn bool
	var rtnRow di.DbRow
	rtnRow.Row = []string{"1", "test2"}
	mdb.MockTestRow = &rtnRow
	// dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var q = "select count(*) from test "
	var a []interface{}
	rowPtr := dbim.Test(q, a...)
	fmt.Println("Mock Records found: ", rowPtr)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		fmt.Print("Mock Records found during test ")
		fmt.Println(int64Val)
		if err2 != nil {
			fmt.Print(err2)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	if !rtn {
		t.Fail()
	}
}

func TestMyDBMock_Insert(t *testing.T) {
	var mdb MyDBMock
	mdb.MockInsertSuccess1 = true
	mdb.MockInsertID1 = 1
	// dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var q = "insert into test (name, address) values(?, ?)"
	var a []interface{}
	a = append(a, "test insert 1", "123 main st")
	suc, id := dbim.Insert(q, a...)
	if !suc || id < 1 {
		t.Fail()
	} else {
		iid1 = id
		fmt.Println("mock inserted id: ", id)
	}
}

func TestMyDBMock_Insert2(t *testing.T) {
	var mdb MyDBMock
	mdb.MockInsertSuccess1 = true
	mdb.MockInsertID1 = 1
	mdb.MockInsertSuccess2 = true
	mdb.MockInsertID2 = 1
	// dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var q = "insert into test (name, address) values(?, ?)"
	var a []interface{}
	a = append(a, "test insert 1", "123 main st")
	dbim.Insert(q, a...)
	suc, id := dbim.Insert(q, a...)
	if !suc || id < 1 {
		t.Fail()
	} else {
		iid1 = id
		fmt.Println("mock inserted id: ", id)
	}
}

func TestMyDBMock_Insert3(t *testing.T) {
	var mdb MyDBMock
	mdb.MockInsertSuccess1 = true
	mdb.MockInsertID1 = 1
	mdb.MockInsertSuccess2 = true
	mdb.MockInsertID2 = 1
	mdb.MockInsertSuccess3 = true
	mdb.MockInsertID3 = 1
	// dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var q = "insert into test (name, address) values(?, ?)"
	var a []interface{}
	a = append(a, "test insert 1", "123 main st")
	dbim.Insert(q, a...)
	dbim.Insert(q, a...)
	suc, id := dbim.Insert(q, a...)
	if !suc || id < 1 {
		t.Fail()
	} else {
		iid1 = id
		fmt.Println("mock inserted id: ", id)
	}
}

func TestMyDBMock_Insert4(t *testing.T) {
	var mdb MyDBMock
	mdb.MockInsertSuccess1 = true
	mdb.MockInsertID1 = 1
	mdb.MockInsertSuccess2 = true
	mdb.MockInsertID2 = 1
	mdb.MockInsertSuccess3 = true
	mdb.MockInsertID3 = 1
	mdb.MockInsertSuccess4 = true
	mdb.MockInsertID4 = 1
	// dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var q = "insert into test (name, address) values(?, ?)"
	var a []interface{}
	a = append(a, "test insert 1", "123 main st")
	dbim.Insert(q, a...)
	dbim.Insert(q, a...)
	dbim.Insert(q, a...)
	suc, id := dbim.Insert(q, a...)
	if !suc || id < 1 {
		t.Fail()
	} else {
		iid1 = id
		fmt.Println("mock inserted id: ", id)
	}
}

func TestMyDBMock_Update(t *testing.T) {
	var mdb MyDBMock
	mdb.MockUpdateSuccess1 = true
	// dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var q = "update test set name = ? , address = ? where id = ? "
	var a []interface{}
	a = append(a, "test insert 2", "123456 main st", iid1)
	suc := dbim.Update(q, a...)
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Update2(t *testing.T) {
	var mdb MyDBMock
	mdb.MockUpdateSuccess1 = true
	mdb.MockUpdateSuccess2 = true
	// dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var q = "update test set name = ? , address = ? where id = ? "
	var a []interface{}
	a = append(a, "test insert 2", "123456 main st", iid1)
	dbim.Update(q, a...)
	suc := dbim.Update(q, a...)
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Update3(t *testing.T) {
	var mdb MyDBMock
	mdb.MockUpdateSuccess1 = true
	mdb.MockUpdateSuccess2 = true
	mdb.MockUpdateSuccess3 = true
	// dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var q = "update test set name = ? , address = ? where id = ? "
	var a []interface{}
	a = append(a, "test insert 2", "123456 main st", iid1)
	dbim.Update(q, a...)
	dbim.Update(q, a...)
	suc := dbim.Update(q, a...)
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Update4(t *testing.T) {
	var mdb MyDBMock
	mdb.MockUpdateSuccess1 = true
	mdb.MockUpdateSuccess2 = true
	mdb.MockUpdateSuccess3 = true
	mdb.MockUpdateSuccess4 = true
	// dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var q = "update test set name = ? , address = ? where id = ? "
	var a []interface{}
	a = append(a, "test insert 2", "123456 main st", iid1)
	dbim.Update(q, a...)
	dbim.Update(q, a...)
	dbim.Update(q, a...)
	suc := dbim.Update(q, a...)
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Get(t *testing.T) {
	var mdb MyDBMock
	var rtnRow di.DbRow
	rtnRow.Row = []string{"2", "test2"}
	mdb.MockRow1 = &rtnRow
	// dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var rtn bool
	var inint int64 = 2
	var q = "select * from test where id = ? "
	var a []interface{}
	a = append(a, inint)
	rowPtr := dbim.Get(q, a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if inint != int64Val {
			fmt.Print(" Mock Get ")
			fmt.Print(inint)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		} else {
			fmt.Print("Mock found id")
			fmt.Print(" = ")
			fmt.Println(int64Val)
			rtn = true
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
	if !rtn {
		t.Fail()
	}
}

func TestMyDBMock_Get2(t *testing.T) {
	var mdb MyDBMock
	var rtnRow di.DbRow
	rtnRow.Row = []string{"2", "test2"}
	mdb.MockRow1 = &rtnRow
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var rtn bool
	var rtn2 bool
	var inint int64 = 2
	var q = "select * from test where id = ? "
	var a []interface{}
	a = append(a, inint)
	rowPtr := dbim.Get(q, a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if inint != int64Val {
			fmt.Print(" Mock Get ")
			fmt.Print(inint)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		} else {
			fmt.Print("Mock found id")
			fmt.Print(" = ")
			fmt.Println(int64Val)
			rtn = true
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
	mdb.MockRow2 = &rtnRow
	rowPtr2 := dbim.Get(q, a...)
	if rowPtr2 != nil {
		foundRow2 := rowPtr2.Row
		fmt.Print("Get ")
		fmt.Println(foundRow2)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow2[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if inint != int64Val {
			fmt.Print(" Mock Get ")
			fmt.Print(inint)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		} else {
			fmt.Print("Mock found id")
			fmt.Print(" = ")
			fmt.Println(int64Val)
			rtn2 = true
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
	if !rtn {
		t.Fail()
	}
	if !rtn2 {
		t.Fail()
	}
}

func TestMyDBMock_Get3(t *testing.T) {
	var mdb MyDBMock
	var rtnRow di.DbRow
	rtnRow.Row = []string{"2", "test2"}
	mdb.MockRow1 = &rtnRow
	mdb.MockRow2 = &rtnRow
	mdb.MockRow3 = &rtnRow
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var rtn bool
	var inint int64 = 2
	var q = "select * from test where id = ? "
	var a []interface{}
	a = append(a, inint)
	dbim.Get(q, a...)
	dbim.Get(q, a...)
	rowPtr := dbim.Get(q, a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if inint != int64Val {
			fmt.Print(" Mock Get ")
			fmt.Print(inint)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		} else {
			fmt.Print("Mock found id")
			fmt.Print(" = ")
			fmt.Println(int64Val)
			rtn = true
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
	if !rtn {
		t.Fail()
	}
}

func TestMyDBMock_Get4(t *testing.T) {
	var mdb MyDBMock
	var rtnRow di.DbRow
	rtnRow.Row = []string{"2", "test2"}
	mdb.MockRow1 = &rtnRow
	mdb.MockRow2 = &rtnRow
	mdb.MockRow3 = &rtnRow
	mdb.MockRow4 = &rtnRow
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var rtn bool
	var inint int64 = 2
	var q = "select * from test where id = ? "
	var a []interface{}
	a = append(a, inint)
	dbim.Get(q, a...)
	dbim.Get(q, a...)
	dbim.Get(q, a...)
	rowPtr := dbim.Get(q, a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if inint != int64Val {
			fmt.Print(" Mock Get ")
			fmt.Print(inint)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		} else {
			fmt.Print("Mock found id")
			fmt.Print(" = ")
			fmt.Println(int64Val)
			rtn = true
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
	if !rtn {
		t.Fail()
	}
}

func TestMyDBMock_GetList(t *testing.T) {
	var mdb MyDBMock
	var rtnRows di.DbRows
	var r1 = []string{"1", "test1"}
	var r2 = []string{"2", "test2"}
	var val [][]string
	val = append(val, r1)
	val = append(val, r2)
	rtnRows.Rows = val
	mdb.MockRows1 = &rtnRows
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var suc bool
	var q = "select * from test where address = ? "
	var a []interface{}
	a = append(a, "123456 main st")
	rowsPtr := dbim.GetList(q, a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Println("Mock rows found: ", foundRows)
		//fmt.Println("GetList results: --------------------------")
		fmt.Println("Mock rows found count: ", len(foundRows))
		if len(foundRows) > 0 {
			suc = true
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_GetList2(t *testing.T) {
	var mdb MyDBMock
	var rtnRows di.DbRows
	var r1 = []string{"1", "test1"}
	var r2 = []string{"2", "test2"}
	var val [][]string
	val = append(val, r1)
	val = append(val, r2)
	rtnRows.Rows = val
	mdb.MockRows1 = &rtnRows
	mdb.MockRows2 = &rtnRows
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var suc bool
	var q = "select * from test where address = ? "
	var a []interface{}
	a = append(a, "123456 main st")
	dbim.GetList(q, a...)
	rowsPtr := dbim.GetList(q, a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Println("Mock rows found: ", foundRows)
		//fmt.Println("GetList results: --------------------------")
		fmt.Println("Mock rows found count: ", len(foundRows))
		if len(foundRows) > 0 {
			suc = true
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_GetList3(t *testing.T) {
	var mdb MyDBMock
	var rtnRows di.DbRows
	var r1 = []string{"1", "test1"}
	var r2 = []string{"2", "test2"}
	var val [][]string
	val = append(val, r1)
	val = append(val, r2)
	rtnRows.Rows = val
	mdb.MockRows1 = &rtnRows
	mdb.MockRows2 = &rtnRows
	mdb.MockRows3 = &rtnRows
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var suc bool
	var q = "select * from test where address = ? "
	var a []interface{}
	a = append(a, "123456 main st")
	dbim.GetList(q, a...)
	dbim.GetList(q, a...)
	rowsPtr := dbim.GetList(q, a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Println("Mock rows found: ", foundRows)
		//fmt.Println("GetList results: --------------------------")
		fmt.Println("Mock rows found count: ", len(foundRows))
		if len(foundRows) > 0 {
			suc = true
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_GetList4(t *testing.T) {
	var mdb MyDBMock
	var rtnRows di.DbRows
	var r1 = []string{"1", "test1"}
	var r2 = []string{"2", "test2"}
	var val [][]string
	val = append(val, r1)
	val = append(val, r2)
	rtnRows.Rows = val
	mdb.MockRows1 = &rtnRows
	mdb.MockRows2 = &rtnRows
	mdb.MockRows3 = &rtnRows
	mdb.MockRows4 = &rtnRows
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var suc bool
	var q = "select * from test where address = ? "
	var a []interface{}
	a = append(a, "123456 main st")
	dbim.GetList(q, a...)
	dbim.GetList(q, a...)
	dbim.GetList(q, a...)
	rowsPtr := dbim.GetList(q, a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Println("Mock rows found: ", foundRows)
		//fmt.Println("GetList results: --------------------------")
		fmt.Println("Mock rows found count: ", len(foundRows))
		if len(foundRows) > 0 {
			suc = true
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Delete(t *testing.T) {
	var mdb MyDBMock
	mdb.MockDeleteSuccess1 = true
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var inint int64 = 2
	var q = "delete from test1 where id = ? "
	var a []interface{}
	a = append(a, inint)
	suc := dbim.Delete(q, a...)
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Delete2(t *testing.T) {
	var mdb MyDBMock
	mdb.MockDeleteSuccess1 = true
	mdb.MockDeleteSuccess2 = true
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var inint int64 = 2
	var q = "delete from test1 where id = ? "
	var a []interface{}
	a = append(a, inint)
	dbim.Delete(q, a...)
	suc := dbim.Delete(q, a...)
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Delete3(t *testing.T) {
	var mdb MyDBMock
	mdb.MockDeleteSuccess1 = true
	mdb.MockDeleteSuccess2 = true
	mdb.MockDeleteSuccess3 = true
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var inint int64 = 2
	var q = "delete from test1 where id = ? "
	var a []interface{}
	a = append(a, inint)
	dbim.Delete(q, a...)
	dbim.Delete(q, a...)
	suc := dbim.Delete(q, a...)
	if !suc {
		t.Fail()
	}
}

func TestMyDBMock_Delete4(t *testing.T) {
	var mdb MyDBMock
	mdb.MockDeleteSuccess1 = true
	mdb.MockDeleteSuccess2 = true
	mdb.MockDeleteSuccess3 = true
	mdb.MockDeleteSuccess4 = true
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	var inint int64 = 2
	var q = "delete from test1 where id = ? "
	var a []interface{}
	a = append(a, inint)
	dbim.Delete(q, a...)
	dbim.Delete(q, a...)
	dbim.Delete(q, a...)
	suc := dbim.Delete(q, a...)
	if !suc {
		t.Fail()
	}
}
func TestMyDBMock_Close(t *testing.T) {
	var mdb MyDBMock
	mdb.MockCloseSuccess = true
	//dbim = &mdb
	dbim = mdb.GetNewDatabase()
	suc := dbim.Close()
	if !suc {
		t.Fail()
	}
}

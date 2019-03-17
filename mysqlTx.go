package mysql

import (
	"database/sql"
	"log"
)

//MyDbTx MyDbTx
type MyDbTx struct {
	Tx *sql.Tx
}

//Insert Insert
func (t *MyDbTx) Insert(query string, args ...interface{}) (bool, int64) {
	var success = false
	var id int64 = -1
	//var stmtIns *sql.Stmt
	stmtIns, err := t.Tx.Prepare(query)
	if err != nil {
		log.Println("Insert transaction prepare err:", err.Error())
	} else {
		defer stmtIns.Close()
		res, err := stmtIns.Exec(args...)
		if err != nil {
			log.Println("Insert transaction Exec err:", err.Error())
		} else {
			id, _ = res.LastInsertId()
			success = true
		}
	}
	return success, id
}

//Update Update
func (t *MyDbTx) Update(query string, args ...interface{}) bool {
	var success = false
	stmtIns, err := t.Tx.Prepare(query)
	if err != nil {
		log.Println("Update transaction prepare err:", err.Error())
	} else {
		defer stmtIns.Close()
		_, err := stmtIns.Exec(args...)
		if err != nil {
			log.Println("Update transaction Exec err:", err.Error())
		} else {
			success = true
		}
	}
	return success
}

//Delete Delete
func (t *MyDbTx) Delete(query string, args ...interface{}) bool {
	var success = false
	stmtIns, err := t.Tx.Prepare(query)
	if err != nil {
		log.Println("Delete transaction prepare err:", err.Error())
	} else {
		defer stmtIns.Close()
		_, err := stmtIns.Exec(args...)
		if err != nil {
			log.Println("Delete transaction Exec err:", err.Error())
			//t.Tx.Rollback()
		} else {
			success = true
			//t.Tx.Commit()
		}
	}
	return success
}

//Commit Commit
func (t *MyDbTx) Commit() bool {
	var rtn = false
	err := t.Tx.Commit()
	if err != nil {
		log.Println("Commit transaction err:", err.Error())
	} else {
		rtn = true
	}
	return rtn
}

//Rollback Rollback
func (t *MyDbTx) Rollback() bool {
	var rtn = false
	err := t.Tx.Rollback()
	if err != nil {
		log.Println("Rollback transaction err:", err.Error())
	} else {
		rtn = true
	}
	return rtn
}

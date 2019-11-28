package mysql

import (
	"database/sql"
	di "github.com/Ulbora/dbinterface"
	"log"

	//Required
	_ "github.com/go-sql-driver/mysql"
)

//MyDB MyDB
type MyDB struct {
	Host     string
	User     string
	Password string
	Database string
	db       *sql.DB
	err      error
}

//Connect Connect
func (m *MyDB) Connect() bool {
	var rtn = false
	var conStr = m.User + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Database
	m.db, m.err = sql.Open("mysql", conStr)
	if m.err != nil {
		log.Println("Open Error:", m.err.Error())
	} else {
		m.err = m.db.Ping()
		if m.err != nil {
			log.Println("Ping Error:", m.err.Error())
		} else {
			rtn = true
		}
	}
	return rtn
}

//GetNewDatabase GetNewDatabase
func (m *MyDB) GetNewDatabase() di.Database {
	var db di.Database
	db = m
	return db
}

//BeginTransaction BeginTransaction
func (m *MyDB) BeginTransaction() di.Transaction {
	var trans di.Transaction
	var myTrans MyDbTx
	tx, err := m.db.Begin()
	if err != nil {
		log.Println("Transaction Error:", err.Error())
	} else {
		myTrans.Tx = tx
	}
	trans = &myTrans
	return trans
}

//Test Test
func (m *MyDB) Test(query string, args ...interface{}) *di.DbRow {
	return m.Get(query, args...)
}

//Insert Insert
func (m *MyDB) Insert(query string, args ...interface{}) (bool, int64) {
	var success = false
	var id int64 = -1
	var stmtIns *sql.Stmt
	stmtIns, m.err = m.db.Prepare(query)
	if m.err != nil {
		log.Println("Error:", m.err.Error())
	} else {
		defer stmtIns.Close()
		res, err := stmtIns.Exec(args...)
		if err != nil {
			log.Println("Insert Exec err:", err.Error())
		} else {
			id, err = res.LastInsertId()
			if err != nil {
				log.Println("Error:", err.Error())
			} else {
				success = true
			}
		}
	}
	return success, id
}

//Update Update
func (m *MyDB) Update(query string, args ...interface{}) bool {
	var success = false
	var stmtUp *sql.Stmt
	stmtUp, m.err = m.db.Prepare(query)
	if m.err != nil {
		log.Println("Error:", m.err.Error())
	} else {
		defer stmtUp.Close()
		res, err := stmtUp.Exec(args...)
		if err != nil {
			log.Println("Update Exec err:", err.Error())
		} else {
			log.Println("Update Exec success:")
			affectedRows, err := res.RowsAffected()
			if err != nil && affectedRows == 0 {
				log.Println("Error:", err.Error())
			} else {
				success = true
			}
		}
	}
	return success
}

//Get Get
func (m *MyDB) Get(query string, args ...interface{}) *di.DbRow {
	var rtn di.DbRow
	stmtGet, err := m.db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtGet.Close()
		rows, err := stmtGet.Query(args...)
		defer rows.Close()
		if err != nil {
			log.Println("Get err: ", err)
		} else {
			columns, err := rows.Columns()
			if err != nil {
				log.Println("Error:", err.Error())
			} else {
				rtn.Columns = columns
				rowValues := make([]sql.RawBytes, len(columns))
				scanArgs := make([]interface{}, len(rowValues))
				for i := range rowValues {
					scanArgs[i] = &rowValues[i]
				}
				for rows.Next() {
					err = rows.Scan(scanArgs...)
					if err != nil {
						log.Println("Error:", err.Error())
					}
					for _, col := range rowValues {
						var value string
						if col == nil {
							value = "NULL"
						} else {
							value = string(col)
						}
						rtn.Row = append(rtn.Row, value)
					}
				}
				if err = rows.Err(); err != nil {
					log.Println("Error:", err.Error())
				}
			}
		}
	}
	return &rtn
}

//GetList GetList
func (m *MyDB) GetList(query string, args ...interface{}) *di.DbRows {
	var rtn di.DbRows
	stmtGet, err := m.db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtGet.Close()
		rows, err := stmtGet.Query(args...)
		defer rows.Close()
		if err != nil {
			log.Println("GetList err: ", err)
		} else {
			columns, err := rows.Columns()
			if err != nil {
				log.Println("Error:", err.Error())
			}
			rtn.Columns = columns
			rowValues := make([]sql.RawBytes, len(columns))
			scanArgs := make([]interface{}, len(rowValues))
			for i := range rowValues {
				scanArgs[i] = &rowValues[i]
			}
			for rows.Next() {
				var rowValuesStr []string
				err = rows.Scan(scanArgs...)
				if err != nil {
					log.Println("Error:", err.Error())
				}
				for _, col := range rowValues {
					var value string
					if col == nil {
						value = "NULL"
					} else {
						value = string(col)
					}
					rowValuesStr = append(rowValuesStr, value)
				}
				rtn.Rows = append(rtn.Rows, rowValuesStr)
			}
			if err = rows.Err(); err != nil {
				log.Println("Error:", err.Error())
			}
		}
	}
	return &rtn
}

//Delete Delete
func (m *MyDB) Delete(query string, args ...interface{}) bool {
	var success = false
	var stmt *sql.Stmt
	stmt, m.err = m.db.Prepare(query)
	if m.err != nil {
		log.Println("Error:", m.err.Error())
	} else {
		defer stmt.Close()
		res, err := stmt.Exec(args...)
		if err != nil {
			log.Println("Delete Exec err:", err.Error())
		} else {
			affectedRows, err := res.RowsAffected()
			if err != nil {
				log.Println("Error:", err.Error())
			} else {
				//fmt.Println("Delete Exec success:")
				if affectedRows > 0 {
					success = true
				}
			}
		}
	}
	return success
}

//Close Close
func (m *MyDB) Close() bool {
	var rtn = false
	err := m.db.Close()
	if err != nil {
		log.Println("database close error: ", err)
	} else {
		rtn = true
	}
	return rtn
}

//GO111MODULE=on go mod init github.com/Ulbora/dbinterface_mysql

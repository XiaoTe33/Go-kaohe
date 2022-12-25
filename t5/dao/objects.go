package dao

import (
	"fmt"
	"strconv"
)

func SQLAddObject(oid string, wid string) {
	sqlStr := "insert into objects(oid, o_wid) values( ?, ?)"
	_, err := Db.Exec(sqlStr, oid, wid)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
}

func SQLPutObject(oid string, num int) {
	sqlStr := "select amount from objects where oid = ?"
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		return
	}
	row := stmt.QueryRow(oid)
	a := ""
	err = row.Scan(&a)
	if err != nil {
		return
	}
	i, err := strconv.Atoi(a)
	if err != nil {
		return
	}

	sqlStr = "update objects set amount=? where oid =?"

	_, err = Db.Exec(sqlStr, i+num, oid)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
}
func SQLMoveObject(oid string, wid string) {
	sqlStr := "update objects set o_wid=? where oid =?"

	_, err := Db.Exec(sqlStr, wid, oid)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
}

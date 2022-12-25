package dao

import "fmt"

func SQLAddWarehouse() {
	sqlStr := "insert into warehouses values ()"
	_, err := Db.Exec(sqlStr)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
}

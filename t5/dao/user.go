package dao

import "fmt"

func SQLAddUser(username string, password string) {

	sqlStr := "insert into users(uid, username, password) values(null, ?, ?)"
	_, err := Db.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Println("Exec err")
		return
	}
}

package dao

func SQLIsValidUsername(username string) bool {
	sqlStr := "select username from users where username = ?"
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		return false
	}
	row := stmt.QueryRow(username)
	u := ""
	err = row.Scan(&u)
	if err != nil {
		return true
	}
	return false
}

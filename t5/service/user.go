package service

import "Go-kaohe/t5/dao"

func ServiceAddUser(username string, password string) {
	dao.SQLAddUser(username, password)
}

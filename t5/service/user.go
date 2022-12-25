package service

import "Go-kaohe/t5/dao"

func AddUser(username string, password string) {
	dao.SQLAddUser(username, password)
}

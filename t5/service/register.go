package service

import "Go-kaohe/t5/dao"

func ServiceIsValidUsername(username string) bool {
	return dao.SQLIsValidUsername(username)
}

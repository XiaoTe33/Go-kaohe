package service

import "Go-kaohe/t5/dao"

func IsValidUsername(username string) bool {
	return dao.SQLIsValidUsername(username)
}

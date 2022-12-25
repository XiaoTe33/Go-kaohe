package service

import "Go-kaohe/t5/dao"

func ServiceAddObject(oid string, wid string) {
	dao.SQLAddObject(oid, wid)
}

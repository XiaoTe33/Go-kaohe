package service

import "Go-kaohe/t5/dao"

func AddObject(oid string, wid string) {
	dao.SQLAddObject(oid, wid)
}

func PutObject(oid string, num int) {
	dao.SQLPutObject(oid, num)
}
func TakeAwayObject(oid string, num int) {
	dao.SQLPutObject(oid, num*(-1))
}

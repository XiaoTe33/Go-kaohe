package api

import "github.com/gin-gonic/gin"

func InitRouters() {
	r := gin.Default()

	r.Use(Middleware())

	r.POST("/Register", Register)
	r.POST("/Login", Login)
	r.POST("/AddWarehouse", AddWarehouse)
	r.POST("/AddObject", AddObject)
	r.POST("/PutObject", PutObject)
	r.POST("/TakeAwayObject", TakeAwayObject)
	r.POST("/MoveObject", MoveObject)

	r.Run(":8083")
}

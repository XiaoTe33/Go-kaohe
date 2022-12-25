package api

import (
	"Go-kaohe/t5/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/Login" {
			c.Next()
			return
		}
		token := c.PostForm("token")
		if !utils.IsValidToken(token) {
			c.JSON(http.StatusBadRequest, gin.H{"err": "请先登录"})
			c.Abort()
			return
		}
		fmt.Println("jwt认证成功")
		c.Next()
		return

	}
}

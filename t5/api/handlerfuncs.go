package api

import (
	"Go-kaohe/t5/service"
	"Go-kaohe/t5/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !service.IsValidUsername(username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "注册失败",
		})
		return
	}
	service.AddUser(username, password)

	token := utils.CreateJWT(username)
	c.JSON(http.StatusOK, gin.H{
		"msg":   "注册成功",
		"token": token,
	})

}

func Login(c *gin.Context) {
	token := c.PostForm("token")
	if utils.IsValidToken(token) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "登陆成功",
		})
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !service.IsValidUsernameAndPassword(username, password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "用户名或密码错误",
		})
		return
	}
	token = utils.CreateJWT(username)
	c.JSON(http.StatusOK, gin.H{
		"msg":   "登陆成功",
		"token": token,
	})
	return
}

func AddWarehouse(c *gin.Context) {
	service.AddWarehouse()
	c.JSON(http.StatusOK, gin.H{
		"msg": "仓库添加成功",
	})
}

func AddObject(c *gin.Context) {
	oid := c.PostForm("oid")
	wid := c.PostForm("wid")
	service.AddObject(oid, wid)
	c.JSON(http.StatusOK, gin.H{
		"msg": "货物增加成功",
	})
}

func PutObject(c *gin.Context) {
	oid := c.PostForm("oid")
	amount := c.PostForm("amount")
	num, err := strconv.Atoi(amount)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "货物上架成功",
	})
	service.PutObject(oid, num)
}
func TakeAwayObject(c *gin.Context) {
	oid := c.PostForm("oid")
	amount := c.PostForm("amount")
	num, err := strconv.Atoi(amount)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "货物上架成功",
	})
	service.TakeAwayObject(oid, num)
}

func MoveObject(c *gin.Context) {

}

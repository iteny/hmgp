package admin

import (
	"fmt"
	"hmgp/common"

	"github.com/gin-gonic/gin"
)

type LoginCtrl struct {
	common.BaseCtrl
}
type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type Account struct {
	Id       int64
	Username string
	Password string
}

func LoginCtrlObject() *LoginCtrl {
	return &LoginCtrl{}
}
func (e *LoginCtrl) Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"title": "admin",
	})
}
func (e *LoginCtrl) LoginAjax(c *gin.Context) {

	var form LoginForm
	var account Account
	if c.ShouldBind(&form) == nil {
		fmt.Println(form.Username, form.Password)
		result := e.Sql().Where("username = ? AND password = ?", form.Username, form.Password).Find(&account)
		fmt.Println(result.Error)
		fmt.Println(account)
		if form.Username == "admin" && form.Password == "admin" {
			c.JSON(200, gin.H{"status": 1, "info": "登录成功"})
		} else {
			c.JSON(200, gin.H{"status": 4, "info": "登录失败"})
		}
	}
}

package routers

import (
	"hmgp/ctrl/admin"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello q1mi!",
	})
}

// SetupRouter 配置路由信息
func Admin(e *gin.Engine) {
	login := admin.LoginCtrlObject()
	e.GET("/admin/login", login.IndexHandler)
	e.GET("/hello", helloHandler)

}

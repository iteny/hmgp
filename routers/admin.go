package routers

import (
	"hmgp/ctrl/admin"
	"hmgp/middle"
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
	index := admin.IndexCtrlObject()
	e.GET("/admin/login", login.Login)

	e.POST("/admin/loginPost", login.LoginAjax)
	e.Use(middle.LoginVerify)
	e.GET("/hello", helloHandler)
	e.GET("/admin/index", index.Index)
	e.GET("/admin/home", index.Home)
	e.GET("/admin/menu", index.Menu)
	e.POST("/admin/getMenu", index.GetMenu)
	e.GET("/admin/addMenu/:pid/:type", index.AddMenu)
	e.POST("/admin/addMenuAjax", index.AddMenuAjax)
}

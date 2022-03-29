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
	admin := admin.AdminCtrlObject()
	e.GET("/admin/login", login.Login)

	e.POST("/admin/loginPost", login.LoginAjax)
	e.Use(middle.LoginVerify)
	e.GET("/hello", helloHandler)
	e.GET("/admin/index", admin.Index)
	e.GET("/admin/home", admin.Home)
	e.GET("/admin/menu", admin.Menu)
	e.POST("/admin/getMenu", admin.GetMenu)
	e.GET("/admin/addCatalog/:pid/:type", admin.AddCatalog)
	e.GET("/admin/addMenu/:pid/:name", admin.AddMenu)
	e.GET("/admin/editMenu/:id", admin.EditMenu)
	e.POST("/admin/addMenuAjax", admin.AddMenuAjax)
	e.POST("/admin/editMenuAjax", admin.EditMenuAjax)
	e.POST("/admin/delMenuAjax", admin.DelMenuAjax)
	//河马股票app
	e.GET("/admin/gpCategory", admin.Category)
	e.POST("/admin/getGpCategory", admin.GetCategoryAjax)
}

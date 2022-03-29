package routers

import (
	"hmgp/ctrl/gpapp"
	"hmgp/middle"

	"github.com/gin-gonic/gin"
)

// func helloHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Hello q1mi!",
// 	})
// }

// SetupRouter 配置路由信息
func Category(e *gin.Engine) {
	// login := admin.LoginCtrlObject()
	cg := gpapp.CategoryCtrlObject()
	// e.GET("/admin/login", login.Login)

	// e.POST("/admin/loginPost", login.LoginAjax)
	e.Use(middle.LoginVerify)
	e.GET("/gpapp/category", cg.Category)
	e.POST("/gpapp/getCategory", cg.GetCategoryAjax)
	// e.GET("/hello", helloHandler)
	// e.GET("/admin/index", index.Index)
	// e.GET("/admin/home", index.Home)
	// e.GET("/admin/menu", index.Menu)
	// e.POST("/admin/getMenu", index.GetMenu)
	// e.GET("/admin/addCatalog/:pid/:type", index.AddCatalog)
	// e.GET("/admin/addMenu/:pid/:name", index.AddMenu)
	// e.GET("/admin/editMenu/:id", index.EditMenu)
	// e.POST("/admin/addMenuAjax", index.AddMenuAjax)
	// e.POST("/admin/editMenuAjax", index.EditMenuAjax)
	// e.POST("/admin/delMenuAjax", index.DelMenuAjax)
}

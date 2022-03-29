package admin

import (
	"hmgp/common/sql"

	"github.com/gin-gonic/gin"
)

func (e *AdminCtrl) Category(c *gin.Context) {
	c.HTML(200, "admin/gp_category.html", gin.H{})
}

//获取菜单
func (e *AdminCtrl) GetCategoryAjax(c *gin.Context) {
	var category []sql.GpCategory
	e.Sql().Order("sort").Find(&category)
	// ar := sql.RecursiveMenu(menus, 0)
	c.JSON(200, gin.H{"data": category})
}

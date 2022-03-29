package admin

import (
	"fmt"
	"hmgp/common"
	"hmgp/common/form"
	"hmgp/common/sql"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminCtrl struct {
	common.BaseCtrl
}

func AdminCtrlObject() *AdminCtrl {
	return &AdminCtrl{}
}
func (e *AdminCtrl) Index(c *gin.Context) {
	var menus []sql.Menu
	e.Sql().Order("sort").Find(&menus)
	ar := sql.RecursiveMenu(menus, 0)
	session := sessions.Default(c)
	username := session.Get("username")
	c.HTML(200, "admin/index.html", gin.H{
		"title":    ar,
		"username": username,
	})
}
func (e *AdminCtrl) Home(c *gin.Context) {
	c.HTML(200, "admin/home.html", gin.H{
		"title": "index",
	})
}

//菜单页面
func (e *AdminCtrl) Menu(c *gin.Context) {
	c.HTML(200, "admin/menu.html", gin.H{})
}

//获取菜单
func (e *AdminCtrl) GetMenu(c *gin.Context) {
	var menus []sql.Menu
	e.Sql().Order("sort").Find(&menus)
	ar := sql.RecursiveMenu(menus, 0)
	c.JSON(200, gin.H{"data": ar})
}

//新增菜单页面
func (e *AdminCtrl) AddMenu(c *gin.Context) {
	pid := c.Param("pid")
	name := c.Param("name")
	c.HTML(200, "admin/add_menu.html", gin.H{
		"pid":  pid,
		"name": name,
	})
}

//新增目录页面
func (e *AdminCtrl) AddCatalog(c *gin.Context) {
	pid := c.Param("pid")
	itype := c.Param("type")
	c.HTML(200, "admin/add_catalog.html", gin.H{
		"pid":  pid,
		"type": itype,
	})
}

//编辑菜单页面
func (e *AdminCtrl) EditMenu(c *gin.Context) {
	id := c.Param("id")
	var menu = sql.Menu{}
	e.Sql().First(&menu, id)
	c.HTML(200, "admin/edit_menu.html", gin.H{
		"result": menu,
	})
}

//新增菜单ajax
func (e *AdminCtrl) AddMenuAjax(c *gin.Context) {
	var form form.AddMenuForm
	if c.ShouldBind(&form) == nil {
		menu := sql.Menu{Name: form.Name, Pid: form.Pid, Type: form.Type, Url: form.Url, Sort: form.Sort, Status: form.Status, Icon: form.Icon}
		result := e.Sql().Create(&menu)
		fmt.Println("asdfdsaf", result)
		if result.Error == nil {
			c.JSON(200, gin.H{"status": 1, "info": "新增成功"})
		} else {
			c.JSON(200, gin.H{"status": 2, "info": "新增失败"})
		}
	} else {
		c.JSON(200, gin.H{"status": 3, "info": "数据验证失败"})
	}
}

//修改菜单ajax
func (e *AdminCtrl) EditMenuAjax(c *gin.Context) {
	var form form.EditMenuForm
	if c.ShouldBind(&form) == nil {
		result := e.Sql().Model(&sql.Menu{}).Where("id = ?", form.Id).Updates(map[string]interface{}{"Name": form.Name, "Type": form.Type, "Url": form.Url, "Sort": form.Sort, "Status": form.Status, "Icon": form.Icon})
		if result.Error == nil {
			c.JSON(200, gin.H{"status": 1, "info": "修改成功"})
		} else {
			c.JSON(200, gin.H{"status": 2, "info": "修改失败"})
		}
	} else {
		c.JSON(200, gin.H{"status": 3, "info": "数据验证失败"})
	}
}

//删除菜单
func (e *AdminCtrl) DelMenuAjax(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	pid, _ := strconv.Atoi(c.PostForm("pid"))
	menu := sql.Menu{}
	if pid == 0 {
		e.Sql().Delete(&menu, id)
		e.Sql().Where("pid = ?", id).Delete(sql.Menu{})
	} else {
		e.Sql().Delete(&menu, id)
	}
	c.JSON(200, gin.H{"status": 1, "info": "删除成功"})
}

package admin

import (
	"fmt"
	"hmgp/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IndexCtrl struct {
	common.BaseCtrl
}
type Menu struct {
	// gorm.Model
	Id       int    `json:"id"`
	Url      string `json:"url"`
	Name     string `json:"name"`
	Pid      int    `json:"pid"`
	Status   int    `json:"status"`
	Sort     int    `json:"sort"`
	Icon     string `json:"icon"`
	Type     int    `json:"type"`
	Children []Menu `json:"children" gorm:"foreignKey:pid"`
}
type AddMenuForm struct {
	Pid    string `form:"pid" binding:"required"`
	Type   string `form:"type" binding:"required"`
	Name   string `form:"name" binding:"required"`
	Url    string `form:"menUrl" binding:"required"`
	Sort   string `form:"sort" binding:"required"`
	Status string `form:"status" binding:"required"`
	Icon   string `form:"icon" binding:"required"`
}

func IndexCtrlObject() *IndexCtrl {
	return &IndexCtrl{}
}
func (e *IndexCtrl) Index(c *gin.Context) {
	var menus []Menu
	e.Sql().Order("sort").Find(&menus)
	ar := RecursiveMenu(menus, 0)
	// fmt.Println(ar)
	// c.JSON(200, gin.H{"status": 1, "info": ar})
	c.HTML(200, "admin/index.html", gin.H{
		"title": ar,
	})
}
func (e *IndexCtrl) Home(c *gin.Context) {
	c.HTML(200, "admin/home.html", gin.H{
		"title": "index",
	})
}
func RecursiveMenu(arr []Menu, pid int) (ar []Menu) {
	array := make([]Menu, 0)
	for k, v := range arr {
		if pid == v.Pid {

			// arr[k].Level = level + 1
			// fmt.Println(arr[k])
			array = append(array, arr[k])
			// fmt.Printf("%#v", array)

		}
	}
	for tk, tv := range array {
		rm := RecursiveMenu(arr, tv.Id)
		for sk := range rm {
			array[tk].Children = append(array[tk].Children, rm[sk])
		}

		// array = append(array, rm[km])
		// array[km].Level = array[km].Level + 1
	}
	return array
}

//菜单页面
func (e *IndexCtrl) Menu(c *gin.Context) {
	// var menus []Menu
	// e.Sql().Order("sort").Find(&menus)
	// ar := RecursiveMenu(menus, 0)
	// fmt.Println(ar)
	// c.JSON(200, gin.H{"status": 1, "info": ar})
	c.HTML(200, "admin/menu.html", gin.H{
		// "data": ar,
	})
}

//获取菜单
func (e *IndexCtrl) GetMenu(c *gin.Context) {
	var menus []Menu
	e.Sql().Order("sort").Find(&menus)
	ar := RecursiveMenu(menus, 0)
	// fmt.Println(ar)
	c.JSON(200, gin.H{"data": ar})
}

//新增菜单页面
func (e *IndexCtrl) AddMenu(c *gin.Context) {
	pid := c.Param("pid")
	itype := c.Param("type")
	fmt.Println("sadfsadf" + itype)
	c.HTML(200, "admin/add_menu.html", gin.H{
		"pid":  pid,
		"type": itype,
	})
}

//新增菜单ajax
func (e *IndexCtrl) AddMenuAjax(c *gin.Context) {
	fmt.Println("asdfdsa2222222")
	var form AddMenuForm
	if c.ShouldBind(&form) == nil {
		fmt.Println(form.Pid, form.Type)
		pid, _ := strconv.Atoi(form.Pid)
		itype, _ := strconv.Atoi(form.Type)
		sort, _ := strconv.Atoi(form.Sort)
		status, _ := strconv.Atoi(form.Status)
		menu := Menu{Name: form.Name, Pid: pid, Type: itype, Url: form.Url, Sort: sort, Status: status, Icon: form.Icon}
		fmt.Println(menu)
		result := e.Sql().Create(&menu)
		fmt.Println("asdfdsa111111111111")
		if result.Error == nil {
			c.JSON(200, gin.H{"status": 1, "info": "新增成功"})
		} else {
			c.JSON(200, gin.H{"status": 2, "info": "新增失败"})
		}
	}
}

//删除菜单
func (e *IndexCtrl) DelMenuAjax(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	pid, _ := strconv.Atoi(c.PostForm("pid"))
	// fmt.Println(id, pid)
	menu := Menu{}
	if pid == 0 {
		e.Sql().Delete(&menu, id)
		e.Sql().Where("pid = ?", id).Delete(Menu{})
	} else {
		e.Sql().Delete(&menu, id)
	}
	c.JSON(200, gin.H{"status": 1, "info": "删除成功"})
}

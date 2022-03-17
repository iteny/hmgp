package admin

import (
	"fmt"
	"hmgp/common"

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
	Isshow   int    `json:"isshow"`
	Sort     int    `json:"sort"`
	Icon     string `json:"icon"`
	Children []Menu `json:"children" gorm:"foreignKey:pid"`
}

func IndexCtrlObject() *IndexCtrl {
	return &IndexCtrl{}
}
func (e *IndexCtrl) Index(c *gin.Context) {
	var menus []Menu
	e.Sql().Order("sort").Find(&menus)
	ar := RecursiveMenu(menus, 0)
	fmt.Println(ar)
	// c.JSON(200, gin.H{"status": 1, "info": ar})
	c.HTML(200, "index.html", gin.H{
		"title": ar,
	})
}
func (e *IndexCtrl) Home(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{
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
func (e *IndexCtrl) Menu(c *gin.Context) {
	// var menus []Menu
	// e.Sql().Order("sort").Find(&menus)
	// ar := RecursiveMenu(menus, 0)
	// fmt.Println(ar)
	// c.JSON(200, gin.H{"status": 1, "info": ar})
	c.HTML(200, "menu.html", gin.H{
		"title": "asdfsad",
	})
}

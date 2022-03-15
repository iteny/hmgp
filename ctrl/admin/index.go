package admin

import (
	"hmgp/common"

	"github.com/gin-gonic/gin"
)

type IndexCtrl struct {
	common.BaseCtrl
}

func IndexCtrlObject() *IndexCtrl {
	return &IndexCtrl{}
}
func (e *IndexCtrl) Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "index",
	})
}

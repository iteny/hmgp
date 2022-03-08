package admin

import "github.com/gin-gonic/gin"

type LoginCtrl struct {
}

func LoginCtrlObject() *LoginCtrl {
	return &LoginCtrl{}
}
func (e *LoginCtrl) IndexHandler(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"title": "admin",
	})
}

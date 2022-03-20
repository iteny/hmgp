package middle

import (
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginVerify(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid")
	// fmt.Println(uid)

	uri := c.Request.RequestURI
	// fmt.Println(uri)
	if uri == "/admin/loginPost" {
		fmt.Println("娶你的")
	} else {
		if uid == nil {
			fmt.Fprint(c.Writer, "<script>parent.window.location = '/admin/login';</script>")
		}
	}
}

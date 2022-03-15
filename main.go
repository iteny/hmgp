package main

import (
	"fmt"
	"hmgp/routers"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("itenyzhaobin"))
	r.Use(sessions.Sessions("mysession", store))
	r.LoadHTMLGlob("view/**/*")
	// r.Use(gin.Recovery())
	r.Static("/static", "./static")
	// 加载静态资源
	// r.StaticFS("/static", http.Dir("./static"))
	// 加载模板文件

	// // 加载404错误页面
	// r.NoRoute(func(c *gin.Context) {
	// 	// 实现内部重定向
	// 	c.HTML(http.StatusOK, "404.html", gin.H{
	// 		"title": "404",
	// 	})
	// })
	// 设置500提示中间件
	// r.Use(errorHttp)
	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))

	routers.Admin(r)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()

}

// errorHttp 统一500错误处理函数
// func errorHttp(c *gin.Context) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			//打印错误堆栈信息
// 			log.Printf("panic: %v\n", r)
// 			debug.PrintStack()
// 			//封装通用json返回
// 			c.HTML(200, "500.html", gin.H{
// 				"title": "500",
// 			})
// 		}
// 	}()
// 	c.Next()
// }

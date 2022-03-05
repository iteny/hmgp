package main

import (
	"fmt"
	"hmgp/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.Admin(r)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

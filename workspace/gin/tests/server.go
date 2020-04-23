package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("/ping is called")
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":3000")
}

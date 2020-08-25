package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func costTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		nowTime := time.Now()

		c.Next()

		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, costTime)
	}
}

func main() {
	r := gin.New()
	r.Use(costTime())
	r.GET("/html", func(c *gin.Context) {
		c.JSON(200, "首页")
	})
	r.Run(":8090")
}

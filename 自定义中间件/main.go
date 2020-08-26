package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time", t2)
	}
}

func main() {
	//1.创建路由
	//默认使用两个中间件(Logger(),Recovery())
	r := gin.Default()
	r.Use(MiddleWare())
	// {} 为了代码规范
	{
		r.GET("/ce", func(c *gin.Context) {
			//取值
			req, _ := c.Get("request")
			fmt.Println("request", req)
			c.JSON(200, gin.H{"request": req})
		})
	}
	r.Run(":8080")
}

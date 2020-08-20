package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//get /?wechat=abc    返回abc
		c.String(200, c.Query("wechat"))
		//get 一个没有定义的值的话,返回默认值 123
		c.String(200, c.DefaultQuery("id", "123"))
		// get /map?ids[a]=123&ids[b]=456&ids[c]=789 返回 {"a":"123","b":"456","c":"789"}
		c.JSON(200, c.QueryMap("ids"))

	})
	r.Run(":8080")
}

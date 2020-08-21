package main

import "github.com/gin-gonic/gin"

func main() {
	allUser := []users{{ID: 123, Name: "abc", Age: 22}, {ID: 222, Name: "222", Age: 222}}
	r := gin.Default()
	r.GET("/user/123", func(c *gin.Context) {
		c.JSON(200, users{ID: 123, Name: "zhangsan", Age: 20})
		c.IndentedJSON(200, allUser) //IndentedJSON 用于美化json输出

	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "<b>Hello, world!</b>", //访问json接口 ,返回 {"message":"\u003cb\u003eHello, world!\u003c/b\u003e"}
		})
	})
	r.GET("/pureJson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"message": "<b>Hello, world!</b>", //访问pureJson接口,返回{"message":"<b>Hello, world!</b>"}
		})
	})
	r.Run(":8080")
}

type users struct {
	ID   int    `json:"id"` //定义这个tag 可以使输出的json字段使用tag的名称
	Name string `json:"name"`
	Age  int    `json:"age"`
}

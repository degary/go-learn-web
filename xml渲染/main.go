package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"whchat": "flysonw_org", "blog": "www.flysonw.org"})
	})

	r.GET("/userxml", func(c *gin.Context) {
		c.XML(200, User{ID: 123, Name: "zhangsan", Age: 123})
	})

	r.GET("/mapxml", func(c *gin.Context) {
		allUsers := []User{{ID: 123, Name: "张三", Age: 20}, {ID: 456, Name: "李四", Age: 25}}
		c.XML(200, gin.H{"user123": allUsers})
	})

	r.Run(":8080")

}

type User struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

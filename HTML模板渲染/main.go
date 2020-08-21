package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

//func main() {
//	r := gin.Default()
//	r.LoadHTMLFiles("html/index.html")
//	r.GET("/html", func(c *gin.Context) {
//		c.HTML(200,"index.html","show me the money")
//	})
//	r.Run(":8080")
//}

func main() {
	r := gin.Default()
	r.GET("/html", func(c *gin.Context) {
		c.Status(200)
		const templateText = `微信公众号: {{printf "%s" .}}`
		tmpl, err := template.New("htmlTest").Parse(templateText)
		if err != nil {
			log.Fatalf("parsing: %s", err)
		}
		tmpl.Execute(c.Writer, "flysnow_org")
	})
	r.Run(":8080")
}

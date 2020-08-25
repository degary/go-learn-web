package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"html/template"
)

func MD5(in string) (string, error) {
	hash := md5.Sum([]byte(in))
	return hex.EncodeToString(hash[:]), nil
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"md5": MD5,
	})
	r.LoadHTMLGlob("html/*")
}

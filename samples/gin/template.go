package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()

	//gin框架添加静态文件
	r.Static("/XXX","./static")

	//gin框架添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})

	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Gin Gubao!",
			"link":"<a href='https://www.baidu.com'>咕宝一下 咕宝知道</a>",
		})
	})
	r.GET("/users", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Gin Gubao!",
			"link":"<a href='https://www.google.com'>咕咕一下 鸽鸽知道</a>",
		})
	})
	r.Run(":5555")
}

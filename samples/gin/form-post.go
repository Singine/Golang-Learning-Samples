package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.LoadHTMLFiles("./form-post.html")

	engine.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "form-post.html",nil)
	})


	engine.POST("/login", func(context *gin.Context) {
		//username := context.PostForm("username")
		//password := context.PostForm("password")

		//username := context.DefaultPostForm("username","nobody")
		//password := context.DefaultPostForm("password","nullpassword")

		username,ok := context.GetPostForm("username")
		if !ok {
			username = "nobody"
		}
		password,ok := context.GetPostForm("password")
		if !ok {
			username = "nullpassword"
		}




		context.HTML(http.StatusOK,"form-post.html",gin.H{
			"Username":username,
			"Password":password,
		})

	})




	engine.Run(":5555")

}

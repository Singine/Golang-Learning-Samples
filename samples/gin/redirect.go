package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	//engine.GET("/index", func(context *gin.Context) {
	//	//context.JSON(http.StatusOK, gin.H{
	//	//	"status": "ok",
	//	//})
	//	context.Redirect(http.StatusMovedPermanently, "/login")
	//})

	engine.GET("/redirect", func(context *gin.Context) {
		context.Request.URL.Path = "/destination" //修改请求的地址
		engine.HandleContext(context) //继续后续处理
	})
	engine.GET("/destination", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "done!",
		})
	})

	engine.Run(":5555")
}

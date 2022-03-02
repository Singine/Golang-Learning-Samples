package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method":"GET",
		})
	})
	engine.POST("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method":"POST",
		})
	})
	engine.PUT("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method":"PUT",
		})
	})
	engine.DELETE("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method":"DELETE",
		})
	})

	engine.NoRoute(func(context *gin.Context) {
		//context.HTML(http.StatusNotFound, "form-post.html",gin.H{
		//	"status":"404",
		//})
		//context.Redirect(http.StatusTemporaryRedirect, "/index")
		context.JSON(http.StatusNotFound,gin.H{
			"status":"404",
		})
	})

	
	//engine.GET("/video/index", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"msg":"/video/index",
	//	})
	//})
	//engine.GET("/shop/index", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"msg":"/shop/index",
	//	})
	//})

	//路由组

	//把公共用的前缀提取出来，创建一个路由组
	videoGroup := engine.Group("/video")
	{
		videoGroup.GET("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/1", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "/video/1",
			})
		})
		videoGroup.GET("/2", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "/video/2",
			})
		})
		videoGroup.GET("/3", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "/video/3",
			})
		})
	}

	engine.Run(":5555")
}

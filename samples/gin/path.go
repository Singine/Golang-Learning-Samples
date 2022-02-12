package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//获取请求的URL path 参数
	
	engine := gin.Default()
	
	engine.GET("/:name/:age", func(context *gin.Context) {
		//获取路径参数
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})
	
	engine.Run(":5555")
	
}

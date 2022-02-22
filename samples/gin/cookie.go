package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/cookie/set", func(context *gin.Context) {

		context.SetCookie("name","Gubao",120,"/","localhost",false,true)
		context.JSON(http.StatusOK,gin.H{
			"set cookie": "done!",
		})

	})
	engine.GET("/cookie/get", func(context *gin.Context) {

		v,_ := context.Cookie("name")
		context.JSON(http.StatusOK,gin.H{
			"cookie": v,
		})

	})
	engine.GET("/cookie/clear", func(context *gin.Context) {

		context.SetCookie("name","Gubao",-1,"/","localhost",false,true)
		context.JSON(http.StatusOK,gin.H{
			"clear cookie": "done!",
		})

	})

	engine.Run(":5555")
}

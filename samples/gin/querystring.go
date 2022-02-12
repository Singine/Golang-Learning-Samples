package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/web", func(context *gin.Context) {

		//获取url参数
		//name := context.Query("name")
		//if name == "Gubao" {
		//	context.JSON(http.StatusOK, "BoBo!")
		//}else {
		//	context.JSON(http.StatusOK, "Kuku!")
		//}


		//取不到值就用默认的
		//name := context.DefaultQuery("name","gaga")
		//context.JSON(http.StatusOK, gin.H{
		//	"name": name,
		//})


		//会返回一个布尔值 可以根据其判断来具体操作
		name,ok := context.GetQuery("name")
		if !ok {
			name = "Kuku!"
		}
		context.JSON(http.StatusOK, gin.H{
			"name": name,
		})


	})

	engine.Run(":5555")
}

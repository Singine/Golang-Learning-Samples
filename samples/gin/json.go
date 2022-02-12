package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.GET("/json", func(context *gin.Context) {
		//方法1 使用map
		//data := map[string]interface{}{
		//	"name":"Gubao",
		//	"message":"hi",
		//	"action":"bobo",
		//}
		data1 := gin.H{
			"name":"Gubao",
			"message":"hi",
			"action":"bobo",
		}

		//方法2 使用结构体
		//灵活使用tag 来定制结构体struct的标签首字母小写
		type msg struct{
			Name string `json:"name"`
			Message string `json:"message"`
			Action string `json:"action"`
		}

		data2 := msg{
			"Gubao",
			"hi",
			"bobo",
		}

		context.JSON(http.StatusOK, gin.H{
			"1":data1,
			"2":data2,
		})
	})

	engine.Run(":5555")


}

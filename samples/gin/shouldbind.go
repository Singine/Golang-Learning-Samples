package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {

	engine := gin.Default()

	engine.GET("/", func(context *gin.Context) {
		var u UserInfo
		err := context.ShouldBind(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}else{
			fmt.Printf("打印一下\n%v\n",u)
			context.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	engine.POST("/form", func(context *gin.Context) {
		var u UserInfo
		err := context.ShouldBind(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}else{
			fmt.Printf("打印一下\n%v\n",u)
			context.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	engine.POST("/json", func(context *gin.Context) {
		var u UserInfo
		err := context.ShouldBind(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}else{
			fmt.Printf("打印一下\n%v\n",u)
			context.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	engine.Run(":5555")

}

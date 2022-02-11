package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("url request:",context.FullPath())
		context.Writer.Write([]byte("Hello,Gin\n"))
		//context.JSON(200, gin.H{
		//	"message": "pong",
		//})
	})
	if err:= engine.Run(":8090");err != nil{
		log.Fatal(err.Error())
	}
}
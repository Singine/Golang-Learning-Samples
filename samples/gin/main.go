package main

import "github.com/gin-gonic/gin"

func sayhi(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Hi Gubao!",
	})
}

func main() {
	r := gin.Default()

	r.GET("/gubao", sayhi)

	r.Run(":5555")

}
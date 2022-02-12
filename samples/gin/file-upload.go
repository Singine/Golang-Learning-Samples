package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLFiles("./file-upload.html")
	engine.GET("/upload", func(context *gin.Context) {
		context.HTML(http.StatusOK,"file-upload.html",nil)
	})
	engine.POST("/upload", func(context *gin.Context) {
		file,err := context.FormFile("f1")
		if err != nil {
			context.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
		}else{
			fileDirectory := path.Join("./uploads",file.Filename)
			context.SaveUploadedFile(file, fileDirectory)
			context.JSON(http.StatusOK, gin.H{
				"status": "ok!",
			})
		}
	})
	engine.Run(":5555")
}

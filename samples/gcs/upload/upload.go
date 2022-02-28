package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main()  {
	r := gin.Default()
	r.LoadHTMLFiles("upload.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK,"upload.html",gin.H{
			"message": "Hi Gubao!",
		})
	})
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		str := strings.Split(file.Filename, ".")[1]
		log.Println(file.Filename)
		log.Println(str)
		newfilename := "article_3.png"
		dst := fmt.Sprintf("C:/Users/DZH/go/Golang-Learning-Samples/samples/gcs/upload/files/%s", newfilename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
		//newfilepath := "C:/Users/DZH/go/Golang-Learning-Samples/samples/gcs/upload/files/" + newfilename

		//uploadFile(newfilepath,"gubao_bucket_test",newfilename)
	})

	r.Run(":5555")
}
// uploadFile uploads an object.
func uploadFile(filepath, bucket, object string) error {
	//bucket := "bucket-name"
	//object := "object-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// Open local file.
	f, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object("demo/" + object).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	return nil
}
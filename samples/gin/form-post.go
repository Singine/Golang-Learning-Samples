package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func paramsJsonHandle(ctx *gin.Context) {

	//JSON传值进来
	// 获取request.Body() 中的数据(这里没有进行错误处理)
	// 返回的是字节数组
	dataBytes, _ := ctx.GetRawData()
	// 定义一个map
	var m map[string]interface{}
	// 反序列化 别忘了&
	_ = json.Unmarshal(dataBytes, &m)
	// 数据返回
	ctx.JSON(http.StatusOK, m)
}

func main() {
	engine := gin.Default()

	engine.LoadHTMLFiles("./form-post.html")

	engine.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "form-post.html",nil)
	})

	engine.POST("/login", func(context *gin.Context) {

		//Form表单传值进来
		//username := context.PostForm("username")
		//password := context.PostForm("password")
		//username := context.DefaultPostForm("username","nobody")
		//password := context.DefaultPostForm("password","nullpassword")

		username,ok := context.GetPostForm("username")
		if !ok {
			username = "nobody"
		}
		password,ok := context.GetPostForm("password")
		if !ok {
			username = "nullpassword"
		}

		context.HTML(http.StatusOK,"form-post.html",gin.H{
			"Username":username,
			"Password":password,
		})

	})
	


	engine.Run(":5555")

}

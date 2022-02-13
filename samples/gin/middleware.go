package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// handlerFunc
func demoHandler(context *gin.Context)  {
	context.JSON(http.StatusOK, gin.H{
		"msg":"done!",
	})
	fmt.Printf("----------------\n")
}

// 统计耗时
func coutingTime(context *gin.Context){

	start := time.Now()
	context.Next()  // 继续调用后续的处理函数

	cost := time.Since(start)

	context.JSON(http.StatusOK, gin.H{
		"timeCost": cost,
	})
	fmt.Printf("%v\n",cost)

}

func m1(context *gin.Context)  {
	fmt.Println("m1 in............")
	context.Next()
	fmt.Println("m1 out............")
	value, ok := context.Get("value") //跨中间件获取用get
	if !ok{
		value = "nobody"
	}
	context.JSON(http.StatusOK, gin.H{
		"value": value,
	})
}
func m2(context *gin.Context)  {
	fmt.Println("m2 in............")
	context.Set("value","bobo") //跨中间件传值用set
	context.Next()
	//context.Abort() // 阻止调用后续的处理函数
	//return

	fmt.Println("m2 out............")
}


func authChecking(checking bool)gin.HandlerFunc {
	//连接数据库
	//其他工作
	return func(context *gin.Context) {
		if checking{
			// 判断是否登录用户
			context.Next()
			// if 登录用户
			// else
			// context.Abort()
		}else{
			context.Next()
		}

	}
}


func main() {
	engine := gin.Default()

	engine.Use(m1,m2,authChecking(true)) //全局注册m1m2函数

	engine.GET("/middleware",demoHandler ,coutingTime )

	////路由组中间件 方法1
	//demoGroup1 := engine.Group("/middlewareGroup",authChecking(true))
	//{
	//	demoGroup1.GET("/1", func(context *gin.Context) {})
	//	demoGroup1.GET("/2", func(context *gin.Context) {})
	//	demoGroup1.GET("/3", func(context *gin.Context) {})
	//}
	//
	////路由组中间件 方法2
	//demoGroup2 := engine.Group("/middlewareGroup")
	//demoGroup2.Use(authChecking(true))
	//{
	//	demoGroup2.GET("/1", func(context *gin.Context) {})
	//	demoGroup2.GET("/2", func(context *gin.Context) {})
	//	demoGroup2.GET("/3", func(context *gin.Context) {})
	//}

	engine.Run(":5555")
}

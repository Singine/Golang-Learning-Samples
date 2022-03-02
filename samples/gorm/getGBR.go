package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/getGBR",API_GetReminiscences(true))

	r.Run(":5555")

}

func InitDB(table string) string{
	dsn := "root:Haishigugue33@tcp(149.248.56.113:3306)/" + table + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

type GetReminiscences struct {
	P_title string
	P_location string
	P_tag string
	P_date string
	P_desc string
	P_src string
}
func (GetReminiscences) TableName() string {
	return "pictures"
}

func API_GetReminiscences(status bool)gin.HandlerFunc {
	return func(context *gin.Context) {
		if status{

			db, err := gorm.Open("mysql", InitDB("gubao_storage"))
			if err != nil {
				panic(err)
			}
			defer db.Close()

			var pics []GetReminiscences

			if err := db.Where("p_date LIKE ?", "02/28%").Find(&pics).Error; err != nil {

				context.JSON(http.StatusUnauthorized,gin.H{
					"status":401,
					"error":"无权限调用api",
				})
				context.Abort()
			}else{
				context.JSON(http.StatusOK,gin.H{
					"status":http.StatusOK,
					"data":pics,
				})
				context.Next()
			}


		}else{
			context.JSON(http.StatusMethodNotAllowed,gin.H{
				"status":http.StatusMethodNotAllowed,
				"error":"该api未开放",
			})
			context.Abort()
		}

	}
}

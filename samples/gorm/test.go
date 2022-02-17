package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//定义模型

type Postgiftlsit struct {
	G_id int
	G_name string
	G_date string
	G_location string
}
func (Postgiftlsit) TableName() string {
	return "gift_list"
}




func main() {
	db, err := gorm.Open("mysql", "root:Haishigugue33@(localhost:3306)/gubao_gift?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//把模型和数据库里的表对应起来

	//db.AutoMigrate(&Giftlist{})

	//t := Postgiftlsit{
	//	G_name: "bobo7",
	//	G_date: "2000-04-01",
	//	G_location: "baowo",
	//}

	var g Postgiftlsit

	//db.NewRecord(&g) //判断主键是否为空

	//db.Debug().Create(&g)
	//db.Debug().Find(&g)
	//db.Find(&g, "g_name = ?", "test")
	//db.Model(&g).Create(&t)

	db.Where("g_id=?","2").Delete(&g)

}
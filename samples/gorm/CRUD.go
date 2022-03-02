package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//定义模型

type Giftlist struct {
	Name string
	Date *string `gorm:"default:'2012-08-21'"`
	Location string
}


func main() {
	db, err := gorm.Open("mysql", "root:Haishigugue33@(localhost:3306)/gubao_gift?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//把模型和数据库里的表对应起来

	db.AutoMigrate(&Giftlist{})

	g := Giftlist{
		Name: "bobo5",
		Date: new(string),
		Location: "baowo",
	}


	//db.NewRecord(&g) //判断主键是否为空

	db.Debug().Create(&g)

}
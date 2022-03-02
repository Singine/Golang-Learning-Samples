package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Gift struct {
	Name string
	Date string
	Location string
}

func main() {
	db, err := gorm.Open("mysql", "root:Haishigugue33@(localhost:3306)/gubao_gift?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表 自动迁移

	//db.AutoMigrate(&Gift{})


	//创建数据行

	//u1 := Gift{
	//	"name1",
	//	"2000-01-09",
	//	"baowo",
	//}
	//db.Create(&u1)



	//查询

	var g Gift
	db.First(&g)  //查询第一条数据到g
	fmt.Printf("%v\n",g)


	//更新

	db.Model(&g).Update("Name","boboni")



	//删除数据

	db.Delete(&g)



}
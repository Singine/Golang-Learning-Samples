package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Gift_list2 struct {
	U_id int `gorm:"primary_key;AUTO_INCREMENT;"`
	Name string
	Age sql.NullInt16
	Birthday *time.Time
	Email string `gorm:"type:varchar(100);unique_index"`
	Role string `gorm:"size:255"`
	MemberNum *string `gorm:"unique;not null"`
	Num int `gorm:"AUTO_INCREMENT;"`
	Address string `gorm:"index:addr"`
	IgnoreMe int `gorm:"-"`
}

func (Gift_list2) TableName() string {
	return "Gift_list2"
}

func main() {
	db, err := gorm.Open("mysql", "root:Haishigugue33@(localhost:3306)/gubao_gift?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Gift_list2{})

	db.Table("bobo").CreateTable(&Gift_list2{})

	var deleted_users []Gift_list2
	db.Table("Gift_list2").Find(&deleted_users)
	//// SELECT * FROM deleted_users;

	db.Table("Gift_list2").Where("name = ?", "bobo").Delete(&deleted_users)
	//// DELETE FROM deleted_users WHERE name = 'jinzhu';

}

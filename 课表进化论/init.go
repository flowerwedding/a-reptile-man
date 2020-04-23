package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/dome7?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
	if !db.HasTable(&Student{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Student{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&Class{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Class{}).Error; err != nil {
			panic(err)
		}
	}
	db.Model(&Class{}).AddIndex("idx_Dacourse_num", "Dacourse_num")
}

type Student struct {
	Stunum    int    `gorm:"type:int(255);not null;primary_key"`
	Username  string `gorm:"type:varchar(256);not null;"`
	Nowweek   string `gorm:"type:varchar(256);not null;"`
	Date      string `gorm:"type:varchar(1024);not null;"`
}

type Class struct {
	Dacourse_num string `gorm:"type:varchar(256);not null;primary_key"`
	Course_num   string `gorm:"type:varchar(256);not null;"`
	Day          string `gorm:"type:varchar(256);not null;"`
	Lesson       string `gorm:"type:varchar(256);not null;"`
	Course       string `gorm:"type:varchar(256);not null;"`
	Teacher      string `gorm:"type:varchar(256);not null;"`
	Classroom    string `gorm:"type:varchar(256);not null;"`
	Rewweek      string `gorm:"type:varchar(256);not null;"`
	Type         string `gorm:"type:varchar(256);not null;"`
}

type NewClass struct {//要json输出的结构体要首字母大写，直接println可小写，结构体里顺序决定json输出顺序
	Hash_day     int
	Hash_lesson  int
	Begin_lesson int
	Class
	WeekModel    string
	Weekbegin    int
	Weekend      int
	Period       string
	Week         []int
}
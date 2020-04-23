package main

import (
	"fmt"
)

func insertstudent(student Student){
	if err :=  db.Model(&Student{}).Create(&student).Error; err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(student)
}

func queryclass1(dacourse_num string)Class{
	var class Class
	db.Model(&Class{}).Where("Dacourse_num = ?",dacourse_num).First(&class)
	return class
}

func insertclass(class Class) {
	if err :=  db.Model(&Class{}).Create(&class).Error; err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(class)
}

func querystudent(stunum string)Student{
	var student Student
	db.Model(&Student{}).First(&student,stunum)
	return student
}

func queryclass2(dacourse_num string)Class{//这个函数只能查int型，string不行
	var class Class
	db.Model(&Class{}).First(&class,"'"+dacourse_num+"'")
	return class
}
package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func paxinxicunshujuku(){
	for xuehao:=2019210001;xuehao<=2019215203;xuehao++ {
		var students Student
		body:=paqu(xuehao)
		name,week :=nameandweek(body)

		var date string
		var classes []Class
		if name[10:] != ""{
			date,classes= classxinxi(body)
		}

		students=Student{
			Stunum : xuehao,
			Username: name[10:],
			Nowweek : week,
			Date : date,
		}

		insertstudent(students)
		for _,class := range classes{
			t := queryclass1(class.Dacourse_num)
			if t.Dacourse_num == ""{
				insertclass(class)
			}
		}
	}
}

func chaxunshuju(c *gin.Context){
	xuehao:=c.PostForm("xuehao")
	student := querystudent(xuehao)
	var newclass []NewClass
	arr := strings.Split(student.Date, " ")//获得学生所有的课程编号
	for _,v := range arr{//每门课程的信息获取，以及一周多次课重新存储结构体
		if v != ""{
			classes := queryclass1(v)
			newclass = fenlitongjieke(classes,newclass)
		}
	}
	c.JSON(200,gin.H{"version":"2020.2.15","stuNum":xuehao,"nowWeek":student.Nowweek,"success":"true","status":200,"data":newclass})
}
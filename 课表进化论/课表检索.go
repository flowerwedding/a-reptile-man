package main

import (
	"regexp"
	"strconv"
	"strings"
)

func zhunbeishuchu(class Class)NewClass{
	day, _ := strconv.Atoi(class.Day)
	theday := "一二三四五六七"[(day-1)*3:day*3]

	lesson := strings.Split(class.Lesson, "-")
	lesson1,_ := strconv.Atoi(lesson[0])
	lesson2,_ := strconv.Atoi(lesson[1])
	var thelesson string
	for i := lesson1;i <= lesson2;i++{
		if i > 10 {
			thelesson = thelesson+"十"+"一二"[(i-11)*3:(i-10)*3]
		}else{
			thelesson += "一二三四五六七八九十"[(i-1)*3:i*3]
		}
	}

	if class.Rewweek[:1] == " "{
		class.Rewweek = class.Rewweek[1:]
	}
	rawweek := strings.Split(class.Rewweek, ",")
	var week []int
	for _,u := range rawweek{
		if strings.Index(u,"-") == -1{
			imgReg2:=regexp.MustCompile(`(.*?)周`)
			imgInfo2 := imgReg2.FindStringSubmatch(u)
			i,_ := strconv.Atoi(imgInfo2[1])
			week = append(week,i)
		}else{
			imgReg:=regexp.MustCompile(`(.*?)-(.*?)周`)
			imgInfo := imgReg.FindStringSubmatch(u)

			rawweek1,_ := strconv.Atoi(imgInfo[1])
			rawweek2,_ := strconv.Atoi(imgInfo[2])
			for i := rawweek1; i <=rawweek2 ;i++{
				week = append(week,i)
			}
		}
	}

	var weekModel string
	if strings.Contains(class.Rewweek, "双周") {
		weekModel = "double"
	}else if strings.Contains(class.Rewweek, "双周"){
		weekModel = "singel"
	}else{
		weekModel = "all"
	}

	newclass := NewClass{
		Hash_day     : day-1,
		Hash_lesson  : 0,
		Begin_lesson : lesson1,
		Class        : class,
		WeekModel    : weekModel,
		Weekbegin    : week[0],
		Weekend      : week[len(week)-1],
		Period       : "2",
		Week         : week,
	}
	newclass.Class.Day = "星期"+theday
	newclass.Class.Lesson = thelesson +"节"
	return newclass
}

func fenlitongjieke(classes Class,newclass []NewClass)[]NewClass{
	arr2 := strings.Split(classes.Day, " # ")
	arr3 := strings.Split(classes.Lesson, " # ")
	arr4 := strings.Split(classes.Teacher, " # ")
	arr5 := strings.Split(classes.Classroom, " # ")
	arr6 := strings.Split(classes.Rewweek, " # ")

	for i2,_ := range arr2{
		class := classes
		class.Day = arr2[i2]
		class.Lesson = arr3[i2]
		class.Teacher = arr4[i2]
		class.Classroom = arr5[i2]
		class.Rewweek = arr6[i2]
		newclass = append(newclass,zhunbeishuchu(class))
	}

	return newclass
}
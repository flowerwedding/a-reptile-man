package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func paqu(xuehao int) []byte {
	resp, err := http.Get("http://jwc.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + strconv.Itoa(xuehao))
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	return body
}

func nameandweek(body []byte)(string,string) {
	reBody := strings.ReplaceAll(string(body), "\n", "")
	olReg := regexp.MustCompile(`<div id="head" >(.*?)<div id="logo`)
	olList := olReg.FindAllStringSubmatch(reBody, -1)

	liReg := regexp.MustCompile(`<li>〉〉2019-2020学年2学期 学生课表>>(.*?)</li>`)
	name := liReg.FindStringSubmatch(olList[0][0])

	liReg2 := regexp.MustCompile(`今天是第(.*?)周 星期`)
	week := liReg2.FindStringSubmatch(olList[0][0])

	return name[1], week[1]
}

func classxinxi(body []byte)(string,[]Class){
	var data string
	var flag = 0

	reBody := strings.ReplaceAll(string(body), "\r\n", "")
	olReg := regexp.MustCompile(`<td>备注</td>(.*?)id="kbStuTabs-ttk"`)
	olList := olReg.FindAllStringSubmatch(reBody, -1)

	liReg:=regexp.MustCompile(`<tr >(.*?)</tr>`)
	liList:=liReg.FindAllString(olList[0][0],-1)

	liReg2:=regexp.MustCompile(`<tr>(.*?)</tr>`)
	liList2:=liReg2.FindAllString(olList[0][0],-1)

	classes:=[]Class{}
	for _,v :=range liList{
		liReglimian:=regexp.MustCompile(`<td(.*?)</td>`)
		liListlimian:=liReglimian.FindAllString(v,-1)

		imgReg:=regexp.MustCompile(`<td rowspan='(.*?)'>(.*?)-(.*?)</td>`)
		imgInfo := imgReg.FindStringSubmatch(liListlimian[1])

		imgReg2:=regexp.MustCompile(`<td rowspan='(.*?)'>(.*?)</td>`)
		imgInfo2 := imgReg2.FindStringSubmatch(liListlimian[2])

		imgReg3:=regexp.MustCompile(`<td rowspan='(.*?)'>(.*?)</td>`)
		imgInfo3 := imgReg3.FindStringSubmatch(liListlimian[3])

		imgReg4:=regexp.MustCompile(`<td>(.*?)</td>`)
		imgInfo4 := imgReg4.FindStringSubmatch(liListlimian[5])

		imgReg5:=regexp.MustCompile(`<td>星期(.*?)第(.*?)节(.*?)</td>`)
		imgInfo5 := imgReg5.FindStringSubmatch(liListlimian[6])
		if imgInfo5 == nil {
			imgInfo5 = []string{"", "0", "无安排", "无安排"}
		}

		imgReg6:=regexp.MustCompile(`<td>(.*?)</td>`)
		imgInfo6 := imgReg6.FindStringSubmatch(liListlimian[7])

		t,_ := strconv.Atoi(imgInfo[1])
		for i:=1 ; i < t ;i++{
			liReglimian2:=regexp.MustCompile(`<td(.*?)</td>`)
			liListlimian2:=liReglimian2.FindAllString(liList2[flag],-1)

			imgReg7:=regexp.MustCompile(`<td >(.*?)</td>`)
			imgInfo7 := imgReg7.FindStringSubmatch(liListlimian2[0])
			imgInfo4[1] = imgInfo4[1] + " # " + imgInfo7[1]

			imgReg7 =regexp.MustCompile(`<td>星期(.*?) 第(.*?)节(.*?)</td>`)
			imgInfo7 = imgReg7.FindStringSubmatch(liListlimian2[1])
			imgInfo5[1] = imgInfo5[1] + " # " + imgInfo7[1]
			imgInfo5[2] = imgInfo5[2] + " # " + imgInfo7[2]
			imgInfo5[3] = imgInfo5[3] + " # " + imgInfo7[3]

			imgReg7 =regexp.MustCompile(`<td>(.*?)</td>`)
			imgInfo7 = imgReg7.FindStringSubmatch(liListlimian2[2])
			imgInfo6[1] = imgInfo6[1] + " # " + imgInfo7[1]

			flag++
		}

		classes=append(classes,Class{
			Dacourse_num:  imgInfo2[2],
			Course_num  :  imgInfo[2],
			Day         :  imgInfo5[1],
			Lesson      :  imgInfo5[2],
			Course      :  imgInfo[3],
			Teacher     :  imgInfo4[1],
			Classroom   :  imgInfo6[1],
			Rewweek     :  imgInfo5[3],
			Type        :  imgInfo3[2],
		})
		data = data + imgInfo2[2] + " "//到时候用空格截取
	}

	return data, classes
}
package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func obtain(i int) [] byte{
	userAgent:=`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36`
	//浏览器代理 https://blog.csdn.net/luanxiyuan/article/details/99570457
	url:="https://movie.douban.com/top250?start="+strconv.Itoa(i)+"&filter="
	c := &http.Client{Transport:&http.Transport{TLSClientConfig:&tls.Config{InsecureSkipVerify:true}}} //在http客户端的tls配置中指定证书

	req,err:=http.NewRequest("GET",url,nil)
	//golang中使用http请求 https://www.cnblogs.com/nyist-xsk/p/10550812.html 添加请求头和请求参数
	req.Header.Add("User-Agent",userAgent)//反爬虫，伪装成浏览器
	resp,err:=c.Do(req)//发送的是https请求，所以需要跳过TLS证书验证
	if err!=nil{
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()//获取响应需要注意关闭响应body
	if resp.StatusCode!=200{
		log.Println("Failed to get the website information")
		return nil
	}
	body,err:=ioutil.ReadAll(resp.Body)//string(body)就是网页源代码
	if err!=nil{
		log.Println(err)
		return nil
	}
	return body
}

func deal(body []byte) []Movie {
	reBody:=strings.ReplaceAll(string(body),"\n","")//把所有换行符替换成空格
	olReg:=regexp.MustCompile(`<ol class="grid_view">(.*?)</ol>`)//判断正则表达式是否合法，不合法抛出异常
	olList:=olReg.FindAllStringSubmatch(reBody,-1)//查找所有匹配项

	liReg:=regexp.MustCompile(`<li>(.*?)</li>`)
	liList:=liReg.FindAllString(olList[0][0],-1)

	movies:=[]Movie{}
	for _,v:=range liList{//对每个电影循环，找信息
		imgReg:=regexp.MustCompile(`<img width="\d+" alt="(.*?)" src="(.*?)" class="">`)
		imgInfo:=imgReg.FindStringSubmatch(v)

		imgReg2:=regexp.MustCompile(`导演: (.*?)&`)
		imgInfo2:=imgReg2.FindStringSubmatch(v)

		imgReg3:=regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
		imgInfo3:=imgReg3.FindStringSubmatch(v)

		imgReg4:=regexp.MustCompile(`<span class="inq">(.*?)</span>`)
		imgInfo4:=imgReg4.FindStringSubmatch(v)
        if imgInfo4 == nil {
			imgInfo4 =[] string { "nil","nil" }
		}

		movies=append(movies,Movie{
			Img:  imgInfo[2],
			Name: imgInfo[1],
			director:imgInfo2[1],
			evaluate:imgInfo3[1],
			comment: imgInfo4[1],
		})
	}
	return movies
}

type Movie struct{
	Img   string
	Name  string
	director string
	evaluate string
	comment string
}
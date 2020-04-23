package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	//paxinxicunshujuku()

	router:=gin.Default()
	router.Use(cors.Default())

	router.POST("/KeBiao/xuehao",chaxunshuju)

	_ = router.Run(":8080")
}
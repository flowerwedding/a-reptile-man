package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router:=gin.Default()
	router.Use(cors.Default())

	router.POST("/DouBan/TOP250",func(c *gin.Context){
		var movieSlice []Movie
		for i := 0;i < 10;i++{
			body := obtain(i*25)
			for _,v:=range deal(body){
				movieSlice = append(movieSlice,v)
			}
		}
		c.JSON(200,JsonNested(movieSlice))
	})

	_ = router.Run(":8080")
}


func JsonNested(movieSlice []Movie) []gin.H {
	var movieJsons []gin.H
	var movieJson gin.H
	for _, movies := range movieSlice {
		movieJson = gin.H{
			"picture" : movies.Img,
			"Name" : movies.Name,
			"director" : movies.director,
			"evaluate" :movies.evaluate,
			"comment" :movies.comment,
		}
		movieJsons = append(movieJsons,movieJson)
	}
	return movieJsons
}
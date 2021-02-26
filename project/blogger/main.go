package main

import (
	"github.com/gin-gonic/gin"
	"yytGithub/project/blogger/controller"
	"yytGithub/project/blogger/dal/db"
)

func main() {
	route := gin.Default()

	dns := "root:123456@tcp(localhost:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}

	route.LoadHTMLGlob("project/blogger/views/*")
	route.Static("/static/", "project/blogger/static")
	route.GET("/", controller.IndexHandler)
	route.GET("/article/detail",controller.ArticleDetailHandler)
	_ = route.Run(":8080")
}


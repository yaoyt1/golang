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
	route.GET("/about/me",controller.AboutMeHandler)
	route.Static("/static/", "project/blogger/static")
	route.GET("/", controller.IndexHandler)
	route.GET("/article/detail",controller.ArticleDetailHandler)
	route.GET("/article/new",controller.NewArticleHandler)
	route.POST("/article/submit/",controller.InserArticleHandler)
	route.GET("/leave/index/",controller.LeaveIndexHandler)
	route.POST("/leave/submit/",controller.InserLeaveHandler)
	
	_ = route.Run(":8080")
}


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
	route.GET("/about/me",controller.AboutMeHandler)
	
	//文章详情页面
	route.GET("/article/detail",controller.ArticleDetailHandler)

	//发布文章页面
	route.GET("/article/new",controller.NewArticleHandler)
	
	//发布文章
	route.POST("/article/submit/",controller.InserArticleHandler)

	//留言首页
	route.GET("/leave/index/",controller.LeaveIndexHandler)

	//提交留言
	route.POST("/leave/submit/",controller.InserLeaveHandler)

	//评论提交
	route.POST("/comment/submit/",controller.InserCommentHandler)
	
	//文章分类列表
	route.GET("/category",controller.CategoryArticleHandler)

	_ = route.Run(":8080")
}


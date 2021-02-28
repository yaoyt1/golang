package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"yytGithub/project/blogger/controller"
	"yytGithub/project/blogger/dal/db"
)

func main() {
	route := gin.Default()

	err := db.Init()
	if err != nil {
		panic(err)
	}

	//性能优化
	//[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
	//[GIN-debug] GET    /debug/pprof/             --> github.com/DeanThompson/ginpprof.IndexHandler.func1 (3 handlers)
	//[GIN-debug] GET    /debug/pprof/heap         --> github.com/DeanThompson/ginpprof.HeapHandler.func1 (3 handlers)
	//[GIN-debug] GET    /debug/pprof/goroutine    --> github.com/DeanThompson/ginpprof.GoroutineHandler.func1 (3 handlers)
	//[GIN-debug] GET    /debug/pprof/block        --> github.com/DeanThompson/ginpprof.BlockHandler.func1 (3 handlers)
	//[GIN-debug] GET    /debug/pprof/threadcreate --> github.com/DeanThompson/ginpprof.ThreadCreateHandler.func1 (3 handlers)
	//[GIN-debug] GET    /debug/pprof/cmdline      --> github.com/DeanThompson/ginpprof.CmdlineHandler.func1 (3 handlers)
	//[GIN-debug] GET    /debug/pprof/profile      --> github.com/DeanThompson/ginpprof.ProfileHandler.func1 (3 handlers)
	//[GIN-debug] GET    /debug/pprof/symbol       --> github.com/DeanThompson/ginpprof.SymbolHandler.func1 (3 handlers)
	//[GIN-debug] POST   /debug/pprof/symbol       --> github.com/DeanThompson/ginpprof.SymbolHandler.func1 (3 handlers)
	//[GIN-debug] GET    /debug/pprof/trace        --> github.com/DeanThompson/ginpprof.TraceHandler.func1 (3 handlers)
	//[GIN-debug] GET    /debug/pprof/mutex        --> github.com/DeanThompson/ginpprof.MutexHandler.func1 (3 handlers)
	ginpprof.Wrap(route)

	//加载模板页面
	route.LoadHTMLGlob("./views/*")

	route.Static("/static/", "./static")
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

	//文件上传接口
	route.POST("/upload/file/", controller.UploadFile)
	
	_ = route.Run("0.0.0.0:9001")
}


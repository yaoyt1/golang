package routers

import (
	"blogger/pkg/setting"
	v1 "blogger/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//使用中间件
	r.Use(gin.Logger(), gin.Recovery())

	//设置gin运行模式 ：
	//	DebugMode = "debug"
	//	// ReleaseMode indicates gin mode is release.
	//	ReleaseMode = "release"
	//	// TestMode indicates gin mode is test.
	//	TestMode = "test
	gin.SetMode(setting.RunMode)

	//注册路由
	registerRouterV1(r)

	return r
}

func registerRouterV1(r *gin.Engine) {
	apiv1 := r.Group("api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.UpdateTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
}

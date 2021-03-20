package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	//	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"key": "value",
		})
	})

	return r
}

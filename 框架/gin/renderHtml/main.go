package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	// 加载当前项目templates目录下的所有模板文件
	r.LoadHTMLGlob("框架/gin/renderHtml/templates/*") //这里的模板文件路径一定要是项目的相对路径

	r.GET("/index", func(c *gin.Context) {
		// 将index.tmpl文件中的title替换为我的标题
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "我的标题",
		})
	})

	r.Run(":8000")
}

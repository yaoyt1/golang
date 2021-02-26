package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("框架/gin/fileServer/static")) //这里静态文件一定要写入相对路径
	r.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	//禁用控制台颜色
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")

	//写入文件，和控制台， 这里multiwriter 可以传入多个writer参数，并返回writer对象
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "hello wrold")
	})

	r.Run(":8080")
}

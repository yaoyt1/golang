package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()

	// Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.
	// By default gin.DefaultWriter = os.Stdout.
	r.Use(gin.Logger())

	// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
	//Recovery返回一个中间件，该中间件可以从任何恐慌中恢复，如果有恐慌，则写入500。
	r.Use(gin.Recovery())

	r.GET("", myb)
}

func MyBenchLogger(ctx *gin.Context){
ctx.JSON(http.StatusOK,"")
}
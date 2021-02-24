package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 开启默认路由引擎
	route := gin.Default()
	route.GET("/login", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "hello word ",
		})
	})

	route.GET("/register", registerHandle)

	//分组路由：一般可用于版本控制
	user := route.Group("/user")
	{
		user.GET("/search/:name/:address", searchHandle)
		user.POST("/add", postFormHandle)
	}

	route.POST("/upload", filePostFormHandle)
	route.POST("/uploads", multiparFilePostFormHandle)
	route.Run(":8080")
}

//registerHandle 获取参数
func registerHandle(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "default")
	pwd := ctx.Query("pwd")
	ctx.JSON(200, gin.H{
		"msg":  "hello word ",
		"name": name,
		"pwd":  pwd,
	})
}

//searchHandle 路径参数
func searchHandle(ctx *gin.Context) {
	name := ctx.Param("name")
	address := ctx.Param("address")
	ctx.JSON(200, gin.H{
		"name":    name,
		"address": address,
	})
}

//postFormHandle 表单提交
func postFormHandle(ctx *gin.Context) {
	name := ctx.PostForm("name")
	address := ctx.PostForm("address")
	ctx.JSON(200, gin.H{
		"type":    "postForm",
		"name":    name,
		"address": address,
	})
}

//filePostFormHandle 文件提交
func filePostFormHandle(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}

	fileName := file.Filename
	log.Printf("上传文件：%s[%d]", fileName, file.Size)

	fileName = fmt.Sprintf("框架/gin/asset/%s", fileName)
	ctx.SaveUploadedFile(file, fileName)

	ctx.JSON(http.StatusOK, fmt.Sprintf("上传成功，上传文件：%s[%d]", fileName, file.Size))
}

//multiparFilePostFormHandle 多个文件提交
func multiparFilePostFormHandle(ctx *gin.Context) {
	multiparFiles, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}

	files := multiparFiles.File["file"]
	for i, file := range files {
		fileName := file.Filename
		log.Printf("上传文件：%s[%d]", fileName, file.Size)

		fileName = fmt.Sprintf("框架/gin/asset/%s_%d", fileName, i)
		ctx.SaveUploadedFile(file, fileName)
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("上传成功，上传%d文件。", len(files)))
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Pwd     string `form:"pwd" json:"pwd" binding:"required"`
	UrlType string `json:"type"`
}

func main() {
	route := gin.Default()
	route.POST("/login", paramaneBind)
	route.POST("/loginJson", paramaneBindJson)
	route.POST("/loginXml", paramaneBindXml)
	route.Run(":8080")
}

//paramaneBind 参数绑定， json 渲染1
func paramaneBind(ctx *gin.Context) {
	var u User
	if err := ctx.ShouldBind(&u); err == nil {
		fmt.Printf("login:%#v\n", u)

		//gin json 渲染1
		ctx.JSON(http.StatusOK, gin.H{
			"type": "form",
			"name": u.Name,
			"pwd":  u.Pwd,
		})

	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
}

//paramaneBindJson 参数绑定接收json，json 渲染2
func paramaneBindJson(ctx *gin.Context) {
	var u User
	if err := ctx.ShouldBindJSON(&u); err == nil {
		fmt.Printf("login:%#v\n", u)
		u.UrlType = "json"
		//gin json 渲染2
		ctx.JSON(http.StatusOK, u)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
}

//paramaneBindXml 参数绑定接收json，xml 渲染
func paramaneBindXml(ctx *gin.Context) {
	var u User
	if err := ctx.ShouldBindJSON(&u); err == nil {
		fmt.Printf("login:%#v\n", u)
		u.UrlType = "json"
		//gin xml 渲染
		ctx.XML(http.StatusOK, u)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
}

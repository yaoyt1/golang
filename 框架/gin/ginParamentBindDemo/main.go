package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `form:"name" json:"name" binding:"required"`
	Pwd  string `form:"pwd" json:"pwd" binding:"required"`
}

func main() {
	route := gin.Default()
	route.POST("/login", paramaneBind)
	route.POST("/loginJson", paramaneBindJson)
	route.Run(":8080")
}

func paramaneBind(ctx *gin.Context) {
	var u User
	if err := ctx.ShouldBind(&u); err == nil {
		fmt.Printf("login:%#v\n", u)
		ctx.JSON(http.StatusOK, gin.H{
			"type": "form",
			"name": u.Name,
			"pwd":  u.Pwd,
		})

	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
}

func paramaneBindJson(ctx *gin.Context) {
	var u User
	if err := ctx.ShouldBindJSON(&u); err == nil {
		fmt.Printf("login:%#v\n", u)
		ctx.JSON(http.StatusOK, gin.H{
			"type": "json",
			"name": u.Name,
			"pwd":  u.Pwd,
		})

	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
}

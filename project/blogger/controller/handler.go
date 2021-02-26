package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yytGithub/project/blogger/logic"
)

func IndexHandler(ctx *gin.Context) {
	articlerecordItems, err := logic.GetAricleList(0, 25)
	if err != nil {
		fmt.Printf("查询文章列表报错：%s\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "views/index.html", articlerecordItems)
}

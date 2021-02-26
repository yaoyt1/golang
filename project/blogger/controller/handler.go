package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yytGithub/project/blogger/logic"
)

//IndexHandler文章首页
func IndexHandler(ctx *gin.Context) {
	articlerecordItems, err := logic.GetAricleList(0, 25)
	if err != nil {
		fmt.Printf("查询文章列表报错：%s\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "views/index.html", articlerecordItems)
}

//ArticleDetailHandler 文章详情
func ArticleDetailHandler(ctx *gin.Context) {
	articleIdstr := ctx.Query("articleId")
	articleId, err := strconv.ParseInt(articleIdstr, 10, 64)
	if err != nil {
		fmt.Printf("无效参数：%s\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	articleDetailItem, err := logic.GetAricleDetail(articleId)
	if err != nil {
		fmt.Printf("查询文章内容报错：%s\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "views/detail.html", articleDetailItem)
}

//NewArticleHandler 新增文章加载时候的
func NewArticleHandler(ctx *gin.Context) {
	categoryItems, err := logic.GetAllCategory()
	if err != nil {
		fmt.Printf("获取分类失败，错误如下：%s\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "views/post_article.html", categoryItems)
}

//InserArticleHandler 新增文章
func InserArticleHandler(ctx *gin.Context) {
	author := ctx.PostForm("author")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	categoryIdStr := ctx.PostForm("categoryId")

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}

	err = logic.InserAricle(author, title, content, categoryId)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, "/")
}

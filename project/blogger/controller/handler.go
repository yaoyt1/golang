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

	//查询文章内容
	articleDetailItem, err := logic.GetAricleDetail(articleId)
	if err != nil {
		fmt.Printf("查询文章内容报错：%s\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//查询相关文章
	aboutArticleitems, err := logic.GetAboutAricle(articleId)
	if err != nil {
		fmt.Printf("查询文章内容报错：%s\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//查询上下文章
	upArticleInof, downArticleInof, err := logic.GetUpDownAricle(articleId)
	if err != nil {
		fmt.Printf("获取上下文章失败：%s\n", err)
	}

	var m = make(map[string]interface{}, 5)
	m["articleDetailItem"] = articleDetailItem
	m["aboutArticleitems"] = aboutArticleitems
	m["upArticleInof"] = upArticleInof
	m["downArticleInof"] = downArticleInof

	ctx.HTML(http.StatusOK, "views/detail.html", m)
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

//AboutMeHandler 关于我
func AboutMeHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "views/about.html", nil)
}

func LeaveIndexHandler(ctx *gin.Context) {
	leaveItems, err := logic.GetAllLeave()
	if err != nil {
		fmt.Printf("查询留言失败%s\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}
	ctx.HTML(http.StatusOK, "views/gbook.html", leaveItems)
}

//InserLeaveHandler 新增留言
func InserLeaveHandler(ctx *gin.Context) {
	comment := ctx.PostForm("comment")
	author := ctx.PostForm("author")
	email := ctx.PostForm("email")

	err := logic.InsertLeave(author, email, comment)
	if err != nil {
		fmt.Printf("新增留言失败：%s\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}
	//ctx.Redirect(http.StatusMovedPermanently, "leave/new")
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"time"
	"yytGithub/project/blogger/logic"
	"yytGithub/project/blogger/util"
)

//IndexHandler文章首页
func IndexHandler(ctx *gin.Context) {
	//查询所有文章
	articlerecordItems, err := logic.GetAricleList(0, 25)
	if err != nil {
		fmt.Printf("查询文章列表报错：%s\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//查询所有分类
	categoryItems, err := logic.GetAllCategory()
	if err != nil {
		fmt.Printf("查询分类报错：%s\n", err)
	}

	//查询排行榜数据
	clickRandkingItems, commentRandkingItems, err := logic.GetRandkingAricle()
	if err != nil {
		fmt.Printf("查询排行榜数据报错：%s\n", err)
	}

	var m = make(map[string]interface{}, 5)
	m["articlerecordItems"] = articlerecordItems
	m["categoryItems"] = categoryItems
	m["clickRandkingItems"] = clickRandkingItems
	m["commentRandkingItems"] = commentRandkingItems

	ctx.HTML(http.StatusOK, "views/index.html", m)
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

	//查询评论列表
	commentItems, err := logic.GetComment(articleId)
	if err != nil {
		fmt.Printf("获取上下文章失败：%s\n", err)
	}

	//查询所有分类
	categoryItems, err := logic.GetAllCategory()
	if err != nil {
		fmt.Printf("查询分类报错：%s\n", err)
	}

	var m = make(map[string]interface{}, 5)
	m["articleDetailItem"] = articleDetailItem
	m["aboutArticleitems"] = aboutArticleitems
	m["upArticleInof"] = upArticleInof
	m["downArticleInof"] = downArticleInof
	m["commentItems"] = commentItems
	m["categoryItems"] = categoryItems

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
	ctx.Redirect(http.StatusMovedPermanently, "/leave/index/")
}

//InserCommentHandler 发表评论
func InserCommentHandler(ctx *gin.Context) {
	userName := ctx.PostForm("author")
	email := ctx.PostForm("email")
	articleIdStr := ctx.PostForm("articleId")
	comment := ctx.PostForm("comment")

	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		fmt.Printf("非法文章Id:%d\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	commentId, err := logic.InsertComment(userName, email, comment, articleId)
	if err != nil {
		fmt.Printf("评论失败:%d\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	fmt.Printf("评论成功：%d\n", commentId)

	url := fmt.Sprintf("/article/detail?articleId=%d", articleId)
	ctx.Redirect(http.StatusMovedPermanently, url)
}

//CategoryArticleHandler 文章分类列表
func CategoryArticleHandler(ctx *gin.Context) {
	categoryIdStr := ctx.Query("categoryId")

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		fmt.Printf("非法文章Id:%d\n", categoryId)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	categoryArticleItem, err := logic.GetArticleByCategoryId(categoryId, 0, 50)
	if err != nil {
		fmt.Printf("查询分类文章报错:%s\n", err)
		ctx.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//查询所有分类
	categoryItems, err := logic.GetAllCategory()
	if err != nil {
		fmt.Printf("查询分类报错：%s\n", err)
	}

	//查询排行榜数据
	clickRandkingItems, commentRandkingItems, err := logic.GetRandkingAricle()
	if err != nil {
		fmt.Printf("查询排行榜数据报错：%s\n", err)
	}

	var m = make(map[string]interface{}, 5)
	m["articlerecordItems"] = categoryArticleItem
	m["categoryItems"] = categoryItems
	m["clickRandkingItems"] = clickRandkingItems
	m["commentRandkingItems"] = commentRandkingItems

	ctx.HTML(http.StatusOK, "views/index.html", m)

}

//UploadFile 上传图片
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println(file.Filename)
	rootPath := util.GetRootDir()
	u2 := time.Now().Unix()
	if err != nil {
		return
	}

	ext := path.Ext(file.Filename)
	url := fmt.Sprintf("/static/upload/%d%s", u2, ext)
	dst := filepath.Join(rootPath, url)
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"uploaded": true,
		"url":      url,
	})
}

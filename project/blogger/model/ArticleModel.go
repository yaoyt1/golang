package model

import "time"

//ArticleInfoModel 文章摘要信息
type ArticleInfoModel struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"categoryid"`
	Title        string    `db:"title"`
	Summary      string    `db:"summary"`
	ViewCount    int32     `db:"viewcount"`
	CommentCount int32     `db:"commentcount"`
	Status       int32     `db:"status"`
	UserName     string    `db:"username"`
	CreateTime   time.Time `db:"createtime"`
	UpdateTime   time.Time `db:"updatetime"`
}

//ArticleDetailModel 文章内容信息
type ArticleDetailModel struct {
	ArticleInfoModel
	Content string `db:"content"`
	CategoryModel
}

//AboutArticleModel 相关文章ID
type AboutArticleModel struct {
	Id    int64  `db:"id"`
	Title string `db:"title"`
}

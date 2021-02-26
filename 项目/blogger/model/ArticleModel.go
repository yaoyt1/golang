package model

import "time"

//ArticleInfoModel 文章摘要信息
type ArticleInfoModel struct {
	Id           int64     `db:"Id"`
	CategoryId   int64     `db:"CategoryId"`
	Title        string    `db:"Title"`
	Summary      string    `db:"Summary"`
	ViewCount    int32     `db:"ViewCount"`
	CommentCount int32     `db:"CommentCount"`
	Status       int32     `db:"Status"`
	UserName     string    `db:"UserName"`
	CreateTime   time.Time `db:"CreateTime"`
	UpdateTime   time.Time `db:"UpdateTime"`
}

//ArticleDetailModel 文章内容信息
type ArticleDetailModel struct {
	ArticleInfoModel
	Content string `db:"Content"`
}

//ArticleRecordModel 文章记录model
type ArticleRecordModel struct {
	ArticleInfoModel
	CategoryModel
}

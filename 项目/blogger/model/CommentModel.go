package model

import "time"

//CommentModel 评论model
type CommentModel struct {
	Id         int64     `db:"Id"`
	Content    string    `db:"Content"`
	UserName   string    `db:"UserName"`
	Status     int32     `db:"Status"`
	CreateTime time.Time `db:"CreateTime"`
	ArticleId  int64     `db:"ArticleID"`
}

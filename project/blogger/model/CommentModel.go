package model

import "time"

//CommentModel 评论model
type CommentModel struct {
	Id         int64     `db:"id"`
	Content    string    `db:"content"`
	UserName   string    `db:"username"`
	Status     int32     `db:"status"`
	CreateTime time.Time `db:"createtime"`
	ArticleId  int64     `db:"articleid"`
	Email      string    `db:"email"`
}

package model

import "time"

//LeaveMsgModel 留言model
type LeaveMsgModel struct {
	Id         int64     `db:"id"`
	UserName   string    `db:"username"`
	Email      string    `db:"email"`
	Content    string    `db:"content"`
	CreateTime time.Time `db:"createtime"`
}

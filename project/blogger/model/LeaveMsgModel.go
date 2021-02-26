package model

import "time"

//LeaveMsgModel 留言model
type LeaveMsgModel struct {
	Id         int64     `db:"Id"`
	UserName   string    `db:"UserName"`
	Email      string    `db:"Email"`
	Content    string    `db:"Content"`
	CreateTime time.Time `db:"CreateTime"`
}

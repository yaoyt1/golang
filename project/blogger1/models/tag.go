package models

import (
	"blogger/pkg/util"
	"fmt"
	"time"
)

//Tag
type Tag struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	CreateBy string `json:"createBy:"`
	CreateDt uint64 `json:"createDt:"`
	UpdateBy string `json:"updateBy"`
	UpdateDt uint64 `json:"updateDt:"`
	State    uint   `json:"state"`
}

var allComun = "id, name, createdt, createby, updatedt, updateby, state"

//GetTags 得到标签列表
func GetTags(pageNum, pageSize int, maps interface{}) (tags []*Tag, err error) {
	sqlStr := fmt.Sprintf("select %s from BlogTag ", allComun)

	//拼接条件
	if m, ok := maps.(map[string]interface{}); ok {
		sqlStr += " where "

		var mapi = 0
		for k := range m {
			if mapi != 0 {
				sqlStr += " and "
			}
			sqlStr += util.GetSqlMapToString(k, m[k])
			mapi++
		}
	}
	sqlStr += " limit ?,? "
	err = db.Select(&tags, sqlStr, pageNum, pageSize)
	return
}

//GetTagTotal 得到标签总数
func GetTagTotal(maps interface{}) (count int, err error) {
	sqlStr := "select count(*) from BlogTag "

	//拼接条件
	if m, ok := maps.(map[string]interface{}); ok {
		sqlStr += " where "

		var mapi = 0
		for k := range m {
			if mapi != 0 {
				sqlStr += " and "
			}
			sqlStr += util.GetSqlMapToString(k, m[k])
			mapi++
		}
	}

	err = db.Get(&count, sqlStr)

	return
}

//ExistTagByName
func ExistTagByName(tagName string) (b bool, err error) {
	var num = 0
	sqlStr := "select 1 from BlogTag where  Name=? "
	err = db.Get(&num, sqlStr, tagName)

	b = false
	if err != nil || num != 0 {
		b = true
	}

	return
}

//AddTag
func AddTag(tagName, createBy string) (addTagId uint64, err error) {
	sqlStr := "insert into BlogTag (name, createdt, createby) VALUE (?,?,?)"
	reslut, err := db.Exec(sqlStr, tagName, time.Now().Unix(), createBy)

	if err != nil {
		rowId, _ := reslut.LastInsertId()
		addTagId = uint64(rowId)
	}

	return
}

//UpdateTag
func UpdateTag(tagId int64, maps interface{}) (err error) {
	sqlStr := "update BlogTag "

	//拼接条件
	if m, ok := maps.(map[string]interface{}); ok {
		sqlStr += " set "

		var mapi = 0
		for k := range m {
			if mapi != 0 {
				sqlStr += " , "
			}
			sqlStr += util.GetSqlMapToString(k, m[k])
			mapi++
		}
	}
	sqlStr += " where id = ?"
	_, err = db.Exec(sqlStr, tagId)
	return
}

//UpdateTagStatue
func UpdateTagStatue(tagId int64, updateBy string, state uint) (err error) {
	sqlStr := "update BlogTag set state =? , updatedt=?, updateby=? where id=?"
	_, err = db.Exec(sqlStr, state, time.Now().Unix(), updateBy, tagId)

	return
}

//DeleteTag
func DeleteTag(tagId int64, updateBy string) (err error) {
	err = UpdateTagStatue(tagId, updateBy, 2)
	return
}

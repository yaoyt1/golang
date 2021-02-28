package db

import "yytGithub/project/blogger/model"

func InsertLeveMsg(userName, email, content string) (insertLeaveId int64, err error) {
	sqlStr := "insert into LeaveMsg ( username, email, content) values (?,?,?)"
	result, err := DB.Exec(sqlStr, userName, email, content)
	if err != nil {
		return
	}
	insertLeaveId, err = result.LastInsertId()
	return
}

func SelectAllLeveMsg() (models []*model.LeaveMsgModel, err error) {
	sqlStr := "select id, username, email, content, createtime from LeaveMsg order by createtime desc "
	err = DB.Select(&models, sqlStr)
	return
}

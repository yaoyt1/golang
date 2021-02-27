package logic

import (
	"fmt"
	"strings"
	"yytGithub/project/blogger/dal/db"
	"yytGithub/project/blogger/model"
	"yytGithub/project/blogger/util"
)

func InsertLeave(userName, email, content string) (err error) {
	if strings.TrimSpace(userName) == "" {
		err = fmt.Errorf("姓名不能为空")
		return
	}

	if strings.TrimSpace(email) == "" {
		err = fmt.Errorf("邮箱地址不能为空")
		return
	}

	if strings.TrimSpace(content) == "" {
		err = fmt.Errorf("内容不能为空")
		return
	}

	if util.VerifyEmailFormat(email) == false {
		err = fmt.Errorf("邮箱格式错误email=%s", email)
		return
	}

	insertId, err := db.InsertLeveMsg(userName, email, content)
	if err != nil {
		fmt.Printf("新增留言失败：%s\n", err)
		return
	}
	fmt.Printf("新增留言成功，留言ID：%d", insertId)
	return
}

func GetAllLeave() (items []*model.LeaveMsgModel, err error) {
	items, err = db.SelectAllLeveMsg()
	if err != nil {
		fmt.Printf("查询留言失败：%s", err)
		return
	}
	fmt.Printf("查询留言成功")
	return
}

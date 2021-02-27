package logic

import (
	"fmt"
	"strings"
	"yytGithub/project/blogger/dal/db"
	"yytGithub/project/blogger/model"
	"yytGithub/project/blogger/util"
)

func GetComment(articleId int64) (commentItems []*model.CommentModel, err error) {
	if articleId < 0 {
		err = fmt.Errorf("非法文章ID：%d", articleId)
		return
	}

	commentItems, err = db.SelectCommentByArticleId(articleId)
	return
}
func InsertComment(userName, email, content string, articleId int64) (commentId int64, err error) {
	if strings.TrimSpace(userName) == "" {
		err = fmt.Errorf("作者不能为空")
		return
	}

	if strings.TrimSpace(content) == "" {
		err = fmt.Errorf("评论内容不能为空")
		return
	}

	if strings.TrimSpace(email) == "" {
		err = fmt.Errorf("邮箱不能为空")
		return
	}

	if articleId < 0 {
		err = fmt.Errorf("非法文章ID：%d", articleId)
		return
	}

	if !util.VerifyEmailFormat(email) {
		err = fmt.Errorf("非法邮箱：%s", email)
		return
	}
	commentId, err = db.InserComment(userName, email, content, articleId)
	return
}

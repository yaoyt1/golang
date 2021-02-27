package db

import "yytGithub/project/blogger/model"

//InserComment
func InserComment(userName, email, content string, articleId int64) (commentId int64, err error) {
	commentId = -1
	sqlStr := "insert into Comment (username, email,  content, articleid) values (?,?,?,?)"
	result, err := DB.Exec(sqlStr, userName, email, content, articleId)
	if err != nil {
		return
	}

	commentId, err = result.LastInsertId()
	return
}

//SelectCommentByArticleId
func SelectCommentByArticleId(articleId int64) (commentItems []*model.CommentModel, err error) {
	sqlStr := "select id, content,email, username, status, createtime, articleid from Comment where  ArticleId=? and Status=1 order by CreateTime desc "
	err = DB.Select(&commentItems, sqlStr, articleId)
	return
}

package db

import "yytGithub/project/blogger/model"

//InserComment
func InserComment(userName, email, content string, articleId int64) (commentId int64, err error) {
	tx, err := DB.Begin()
	if err != nil {
		return
	}

	commentId = -1
	sqlStr := "insert into Comment (username, email,  content, articleid) values (?,?,?,?)"
	result, err := DB.Exec(sqlStr, userName, email, content, articleId)
	if err != nil {
		tx.Rollback()
		return
	}

	//更新评论数
	sqlStr = "update Article set  commentcount=commentcount+1 where Id=?"
	_, err = DB.Exec(sqlStr, articleId)
	if err != nil {
		tx.Rollback()
	}

	commentId, err = result.LastInsertId()
	tx.Commit()
	return
}

//SelectCommentByArticleId
func SelectCommentByArticleId(articleId int64) (commentItems []*model.CommentModel, err error) {
	sqlStr := "select id, content,email, username, status, createtime, articleid from Comment where  ArticleId=? and Status=1 order by CreateTime desc "
	err = DB.Select(&commentItems, sqlStr, articleId)
	return
}

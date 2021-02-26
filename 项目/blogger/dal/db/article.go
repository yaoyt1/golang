package db

import (
	"yytGithub/项目/blogger/model"
)

//installArticle 新增文章
func installArticle(articl *model.ArticleDetailModel) (articleId int64, err error) {
	sqlStr := "insert into Article (categoryid, content, title, summary, username) values (?,?,?,?,?)"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	stmt.Query()
	result, err := stmt.Exec(articl.CategoryId, articl.Content, articl.Title, articl.Summary, articl.UserName)
	if err != nil {
		return 0, err
	}

	articleId, _ = result.LastInsertId()
	return
}

//获取文章分页
func getArticleList(pageNum, pageSize int) (articlInfoItems []*model.ArticleInfoModel, err error) {
	sqlStr := `select id, categoryid, title, summary, viewcount, commentcount, status,
       username, createtime, updatetime from Article order by CreateTime DESC  limit ?,?`

	DB.Select(&articlInfoItems, sqlStr, pageNum, pageSize)
	return
}

//根据分类获取文章列表
func getArticleByCategoryIdList(categoryid, pageNum, pageSize int) (articlInfoItems []*model.ArticleInfoModel, err error) {
	sqlStr := `select id, categoryid, title, summary, viewcount, commentcount, status,
       username, createtime, updatetime from Article where CategoryId=? order by CreateTime DESC  limit ?,?`

	DB.Select(&articlInfoItems, sqlStr, categoryid, pageNum, pageSize)
	return
}

//获取文章详细
func getArticleDetailById(categoryid int64) (articlDetailItem *model.ArticleDetailModel, err error) {
	sqlStr := `select id, categoryid, content, title, summary, viewcount, commentcount, status,
       username, createtime, updatetime from Article where id=? `

	DB.Get(&articlDetailItem, sqlStr, categoryid)
	return
}

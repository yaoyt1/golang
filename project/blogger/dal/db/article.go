package db

import (
	"yytGithub/project/blogger/model"
)

//InstallArticle 新增文章
func InstallArticle(articl *model.ArticleDetailModel) (articleId int64, err error) {
	sqlStr := "insert into Article (categoryid, content, title, summary, username) values (?,?,?,?,?)"
	result, err := DB.Exec(sqlStr, articl.CategoryId, articl.Content, articl.Title, articl.Summary, articl.UserName)
	if err != nil {
		return 0, err
	}

	articleId, _ = result.LastInsertId()
	return
}

//SelectArticleList 获取文章分页
func SelectArticleList(pageNum, pageSize int) (articlInfoItems []*model.ArticleInfoModel, err error) {
	sqlStr := `select id, categoryid, title, summary, viewcount, commentcount, status,
       username, createtime, updatetime from Article order by CreateTime DESC  limit ?,?`
	DB.Select(&articlInfoItems, sqlStr, pageNum, pageSize)
	return
}

//SelectArticleByCategoryIdList 根据分类获取文章列表
func SelectArticleByCategoryIdList(categoryid, pageNum, pageSize int) (articlInfoItems []*model.ArticleInfoModel, err error) {
	sqlStr := `select id, categoryid, title, summary, viewcount, commentcount, status,
       username, createtime, updatetime from Article where CategoryId=? order by CreateTime DESC  limit ?,?`

	DB.Select(&articlInfoItems, sqlStr, categoryid, pageNum, pageSize)
	return
}

//SelectArticleDetailById 获取文章详细
func SelectArticleDetailById(categoryid int64) (articlDetailItem *model.ArticleDetailModel, err error) {
	articlDetailItem = new(model.ArticleDetailModel)
	sqlStr := `select id, categoryid, content, title, summary, viewcount, commentcount, status,
       username, createtime, updatetime from Article where status=1 and id=? `

	err = DB.Get(articlDetailItem, sqlStr, categoryid)
	return
}

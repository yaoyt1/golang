package logic

import (
	"fmt"
	"math"
	"strings"
	"yytGithub/project/blogger/dal/db"
	"yytGithub/project/blogger/model"
)

//GetAricleList 获取文章列表
func GetAricleList(pageNum, pageSize int) (articlerecordItems []*model.ArticleDetailModel, err error) {
	if pageNum < 0 || pageSize <= 0 {
		fmt.Printf("非法参数pageNum[%d],pageSize[%d]", pageNum, pageSize)
		return
	}

	articlInfoItems, err := db.SelectArticleList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("查询文章列表失败：%s\n", err)
		return
	}

	if len(articlInfoItems) == 0 {
		return
	}

	categoryIds := getAricleListCategoryId(articlInfoItems)
	if len(categoryIds) == 0 {
		return
	}

	categoryItems, err := db.SelectCategoryByIds(categoryIds)
	if err != nil {
		fmt.Printf("查询分类失败,错误：%s\n", err)
		return
	}

	//聚合数据
	for _, articlInfoItem := range articlInfoItems {
		articleRecord := &model.ArticleDetailModel{
			ArticleInfoModel: *articlInfoItem,
		}
		for _, categoryItem := range categoryItems {
			if articlInfoItem.CategoryId == categoryItem.CategoryId {
				articleRecord.CategoryModel = *categoryItem
			}
		}
		articlerecordItems = append(articlerecordItems, articleRecord)
	}
	return
}

//getAricleListCategoryId
func getAricleListCategoryId(infoModels []*model.ArticleInfoModel) (ids []int64) {
label:
	for _, item := range infoModels {
		categoryId := item.CategoryId
		for _, id := range ids {
			if categoryId == id {
				continue label
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

//GetAricleDetail 根据文章ID获取文章内容
func GetAricleDetail(id int64) (articleDetail *model.ArticleDetailModel, err error) {
	if id < 0 {
		err = fmt.Errorf("无效的文章Id(%d)", id)
		return
	}

	//获取文章信息
	articleDetail, err = db.SelectArticleDetailById(id)
	if err != nil {
		return
	}

	if articleDetail == nil {
		err = fmt.Errorf("根据文章(%d)未查询文章。", id)
		return
	}

	//获取分类信息
	categoryInfo, err := db.SelectCategoryById(articleDetail.ArticleInfoModel.CategoryId)
	if err != nil {
		return
	}

	articleDetail.CategoryModel = *categoryInfo
	return
}

//InserAricle 新增文章
func InserAricle(author, title, content string, categoryId int64) (err error) {
	if len(strings.TrimSpace(author)) == 0 {
		err = fmt.Errorf("作者不能为空")
		return
	}

	if len(strings.TrimSpace(title)) == 0 {
		err = fmt.Errorf("标题不能为空")
		return
	}

	if len(strings.TrimSpace(content)) == 0 {
		err = fmt.Errorf("内容不能为空")
		return
	}

	if categoryId < 0 {
		err = fmt.Errorf("请选择分类id")
		return
	}

	articledatail := new(model.ArticleDetailModel)
	articledatail.UserName = author
	articledatail.Title = title
	articledatail.ArticleInfoModel.CategoryId = categoryId
	articledatail.Content = content

	contentUtf8 := []rune(content)
	minSummary := int(math.Min(float64(len(contentUtf8)), 128.0))
	articledatail.Summary = string(contentUtf8[:minSummary])

	articleId, err := db.InstallArticle(articledatail)
	fmt.Printf("新增文章成功，文章Id：%d\n", articleId)
	return
}

//GetAboutAricle 查询相关文章
func GetAboutAricle(articleId int64) (aboutArticleItems []*model.AboutArticleModel, err error) {
	aboutArticleItems, err = db.SelectAboutArticleList(articleId, 10)
	return
}

func GetUpDownAricle(articleId int64) (upArticleInof, downArticleInof *model.AboutArticleModel, err error) {
	upArticleInof, downArticleInof, err = db.SelectUpDownArticle(articleId)
	return
}

func GetArticleByCategoryId(categoryid int64, pageNum, pageSize int) (items []*model.ArticleDetailModel, err error) {
	if pageNum < 0 || pageSize <= 0 {
		fmt.Printf("非法参数pageNum[%d],pageSize[%d]", pageNum, pageSize)
		return
	}

	CategoryAricleItems, err := db.SelectArticleByCategoryId(categoryid, pageNum, pageSize)
	if err != nil {
		fmt.Printf("查询分类文章列表失败：%s\n", err)
		return
	}

	if len(CategoryAricleItems) == 0 {
		return
	}

	categoryIds := getAricleListCategoryId(CategoryAricleItems)
	if len(categoryIds) == 0 {
		return
	}

	categoryItems, err := db.SelectCategoryByIds(categoryIds)
	if err != nil {
		fmt.Printf("查询分类失败,错误：%s\n", err)
		return
	}

	//聚合数据
	for _, articlInfoItem := range CategoryAricleItems {
		articleRecord := &model.ArticleDetailModel{
			ArticleInfoModel: *articlInfoItem,
		}
		for _, categoryItem := range categoryItems {
			if articlInfoItem.CategoryId == categoryItem.CategoryId {
				articleRecord.CategoryModel = *categoryItem
			}
		}
		items = append(items, articleRecord)
	}
	return
}

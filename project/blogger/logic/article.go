package logic

import (
	"fmt"
	"yytGithub/project/blogger/dal/db"
	"yytGithub/project/blogger/model"
)

//GetAricleList 获取文章列表
func GetAricleList(pageNum, pageSize int) (articlerecordItems []*model.ArticleRecordModel, err error) {
	if pageNum < 0 || pageSize <= 0 {
		fmt.Printf("非法参数pageNum[%d],pageSize[%d]", pageNum, pageSize)
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
		articleRecord := &model.ArticleRecordModel{
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

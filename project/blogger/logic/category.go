package logic

import (
	"yytGithub/project/blogger/dal/db"
	"yytGithub/project/blogger/model"
)

//GetAllCategory 获取所有分类
func GetAllCategory() (categoryItems []*model.CategoryModel, err error) {
	categoryItems, err = db.SelectAllCategory()
	return
}

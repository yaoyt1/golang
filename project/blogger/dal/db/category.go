package db

import (
	"github.com/jmoiron/sqlx"
	"yytGithub/project/blogger/model"
)

//SelectAllCategory 查询所有的分类
func SelectAllCategory() (categoryItems []*model.CategoryModel, err error) {
	sqlStr := "select  id, categoryname, categoryno from Category"
	err = DB.Select(&categoryItems, sqlStr)
	return
}

func SelectCategoryByIds(ids []int64) (categoryItems []*model.CategoryModel, err error) {
	sqlStr, arg, err := sqlx.In("select  id, categoryname, categoryno from Category where id in (?)", ids)
	if err != nil {
		return
	}
	err = DB.Select(&categoryItems, sqlStr, arg...)
	return
}

func SelectCategoryById(id int64) (category *model.CategoryModel, err error) {
	sqlStr := "select  id, categoryname, categoryno from Category where id =?"
	err = DB.Get(&category, sqlStr, id)
	return
}

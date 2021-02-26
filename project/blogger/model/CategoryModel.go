package model

//CategoryModel 文章分类实体
type CategoryModel struct {
	CategoryId   int64  `db:"id"`
	CategoryName string `db:"categoryname"`
	CategoryNo   int    `db:"categoryno"`
}

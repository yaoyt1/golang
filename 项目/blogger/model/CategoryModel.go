package model

//CategoryModel 文章分类实体
type CategoryModel struct {
	CategoryId   int64  `db:"CategoryId"`
	CategoryName string `db:"CategoryName"`
	CategoryNo   int    `db:"CategoryNo"`
}

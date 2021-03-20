package util

import (
	"blogger/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//GetPage 计算当前页页面大小
func GetPage(ctx *gin.Context) (result int) {
	page, _ := com.StrTo(ctx.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return
}

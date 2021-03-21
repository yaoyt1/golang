package v1

import (
	"blogger/models"
	"blogger/pkg/e"
	"blogger/pkg/setting"
	"blogger/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

//GetTags 获取标签列表
func GetTags(ctx *gin.Context) {
	//创建条件集合
	maps := make(map[string]interface{})

	name := ctx.Query("name")
	if name != "" {
		maps["name"] = name
	}

	state := 1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}
	maps["state"] = state

	//创建返回数据集合
	var err error
	code := e.SUCCESS

	data := make(map[string]interface{})
	data["lists"], err = models.GetTags(util.GetPage(ctx), setting.PageSize, maps)
	if err != nil {
		code = e.ERROR
		goto lable
	}

	data["total"], err = models.GetTagTotal(maps)
	if err != nil {
		code = e.ERROR
	}

lable:
	ctx.JSON(http.StatusOK, e.HttpReturn(code, data))
}

//AddTags 新增标签
func AddTag(ctx *gin.Context) {
	name := ctx.Query("name")
	createBy := ctx.Query("createBy")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空。 ")
	valid.MaxSize(name, 100, "name").Message("名称长度最长只能为100位。 ")
	valid.Required(createBy, "createBy").Message("创建人不能为空。 ")
	valid.MaxSize(createBy, 100, "createBy").Message("创建人长度最长只能为100位。 ")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		goto lable
	}

	if ok, _ := models.ExistTagByName(name); !ok {
		code = e.SUCCESS
		_, err := models.AddTag(name, createBy)
		if err != nil {
			code = e.ERROR_NOT_EXIST_ADD
		}
	} else {
		code = e.ERROR_EXIST_TAG
	}

lable:
	ctx.JSON(http.StatusOK, e.HttpReturn(code, nil))
}

//UpdateTags 更新标签
func UpdateTag(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt64()
	name := ctx.Query("name")
	updateBy := ctx.Query("updateBy")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("参数id不能为空。")
	valid.Required(name, "name").Message("参数标签名称【name】不能为空。")
	valid.Required(updateBy, "createBy").Message("参数更新人【updateBy】不能为空。")
	valid.MaxSize(updateBy, 100, "createBy").Message("参数更新人【updateBy】长度不能超过100。")

	var err error
	code := e.SUCCESS
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		ctx.JSON(http.StatusOK, e.HttpReturn(code, nil))
		return
	}

	var maps = make(map[string]interface{})
	maps["name"] = name
	maps["updateBy"] = updateBy

	err = models.UpdateTag(id, maps)
	if err != nil {
		code = e.ERROR_NOT_EXIST_UPDATE
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, e.HttpReturn(code, nil))
}

//DeleteTags 删除标签
func DeleteTag(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt64()
	updateBy := ctx.Query("updateBy")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("参数id不能为空。")
	valid.Min(id, 1, "id").Message("参数id必须大于0 。")
	valid.Required(updateBy, "updateBy").Message("参数更新人【updateBy】不能为空。")
	valid.MaxSize(updateBy, 100, "updateBy").Message("参数更新人【updateBy】长度不能超过100。")

	var err error
	code := e.SUCCESS
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		ctx.JSON(http.StatusOK, e.HttpReturn(code, nil))
		return
	}

	err = models.DeleteTag(id, updateBy)
	if err != nil {
		code = e.ERROR_NOT_EXIST_UPDATE
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, e.HttpReturn(code, nil))
}

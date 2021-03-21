package e

var codeMsg = map[int]string{
	SUCCESS:                 "ok",
	ERROR:                   "fail",
	INVALID_PARAMS:          "请求参数错误",
	ERROR_EXIST_TAG:         "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:     "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE: "该文章不存在",
	ERROR_NOT_EXIST_ADD:     "新增标签失败",
	ERROR_NOT_EXIST_UPDATE:  "更新标签失败",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

//GetCodeMsg 得到错误消息
func GetCodeMsg(code int) string {
	str, ok := codeMsg[code]
	if ok {
		return str
	}
	return codeMsg[ERROR]
}

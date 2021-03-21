package e

type CodeMsg struct {
	Code    int         `json:"code"`
	CodeMsg string      `json:"codeMsg"`
	Data    interface{} `json:"data"`
}

//HttpReturn http返回方法函数
func HttpReturn(code int, data interface{}) CodeMsg {
	return CodeMsg{
		Code:    code,
		CodeMsg: GetCodeMsg(code),
		Data:    data,
	}
}

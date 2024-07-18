package xerr

var codeText = map[int]string{
	SERVER_COMMON_ERROR: "服务器异常，稍后再尝试",
	REQUEST_PARAM_ERROR: "请求参数有误",
	DB_ERROR:            "数据库繁忙，稍后再尝试",
}

// ErrMsg 错误信息
func ErrMsg(code int) string {
	if msg, ok := codeText[code]; ok {
		return msg
	}
	return codeText[SERVER_COMMON_ERROR]
}

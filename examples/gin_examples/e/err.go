package e

var Msg = map[int]string{
	Success:   "ok",
	Error:     "内部异常",
	BindError: "绑定参数异常",
	Exist:     "已存在",
}

func GetMsg(code int) string {
	if msg, ok := Msg[code]; ok {
		return msg
	}

	return Msg[Error]
}

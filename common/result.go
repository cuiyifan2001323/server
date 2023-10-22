package common

type Result struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Success(data any, code int) Result {
	return Result{
		Msg:  "成功",
		Code: code,
		Data: data,
	}
}

func Err(data string, code int) Result {
	return Result{
		Msg:  "错误",
		Code: code,
		Data: data,
	}
}

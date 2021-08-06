package model

type ResponseBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessData(data interface{}) *ResponseBean {
	return &ResponseBean{200, "", data}
}

func SuccessMsg(msg string) *ResponseBean {
	return &ResponseBean{200, msg, ""}
}

func ErrorFail(code int, errorMsg string) *ResponseBean {
	return &ResponseBean{code, errorMsg, ""}
}

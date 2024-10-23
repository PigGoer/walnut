package exception

import "encoding/json"

// ApiException 是一个自定义异常
// 用于返回给前端的异常信息
// param code 全局自定义异常码
// param message 异常信息
// param httpCode http状态码
type ApiException struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	HttpCode int    `json:"-"`
}

func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		Code:    code,
		Message: message,
	}
}

// Error 实现了error接口
func (a *ApiException) Error() string {
	return a.Message
}

// WithHttpCode 设置http状态码
func (a *ApiException) WithHttpCode(httpCode int) *ApiException {
	a.HttpCode = httpCode
	return a
}

// WithMessage 设置异常信息
func (a *ApiException) WithMessage(msg string) *ApiException {
	a.Message = msg
	return a
}

// String 返回异常信息的json字符串
func (a *ApiException) String() string {
	obj, _ := json.MarshalIndent(a, "", "    ")
	return string(obj)
}

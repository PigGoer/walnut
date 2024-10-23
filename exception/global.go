package exception

import "fmt"

// ErrServerInternal 服务器内部错误
func ErrServerInternal(format string, a ...any) *ApiException {
	return NewApiException(5000,
		fmt.Sprintf(format, a...))
}

// ErrNotFound 未找到资源
func ErrNotFound(format string, a ...any) *ApiException {
	return NewApiException(404,
		fmt.Sprintf(format, a...))
}

// ErrValidateFailed 参数验证失败
func ErrValidateFailed(format string, a ...any) *ApiException {
	return NewApiException(400,
		fmt.Sprintf(format, a...),
	)
}

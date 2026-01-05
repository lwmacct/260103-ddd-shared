package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ============================================================================
// 服务端错误响应函数 (5xx)
// ============================================================================

// InternalError 500 服务器错误
func InternalError(c *gin.Context, details ...any) {
	Failure(c, http.StatusInternalServerError, MsgInternalError, details...)
}

// NotImplemented 501 功能未实现
func NotImplemented(c *gin.Context, message ...string) {
	msg := MsgNotImplemented
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusNotImplemented, msg)
}

// ServiceUnavailable 503 服务不可用
func ServiceUnavailable(c *gin.Context, message ...string) {
	msg := MsgServiceUnavailable
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusServiceUnavailable, msg)
}

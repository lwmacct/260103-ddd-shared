package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ============================================================================
// 客户端错误响应函数 (4xx)
// ============================================================================

// BadRequest 400 请求错误
func BadRequest(c *gin.Context, message string, details ...any) {
	Failure(c, http.StatusBadRequest, message, details...)
}

// ValidationError 400 验证错误
func ValidationError(c *gin.Context, details any) {
	Failure(c, http.StatusBadRequest, MsgValidationFailed, details)
}

// Unauthorized 401 未认证
func Unauthorized(c *gin.Context, message ...string) {
	msg := MsgAuthenticationRequired
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusUnauthorized, msg)
}

// Forbidden 403 无权限
func Forbidden(c *gin.Context, message ...string) {
	msg := MsgAccessForbidden
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusForbidden, msg)
}

// NotFound 404 资源不存在
func NotFound(c *gin.Context, resource string) {
	message := MsgResourceNotFound
	if resource != "" {
		message = fmt.Sprintf(MsgNotFoundFormat, resource)
	}
	Failure(c, http.StatusNotFound, message)
}

// NotFoundMessage 404 资源不存在（自定义消息）
// 用于传递完整的错误消息（如来自 Domain 层的 err.Error()）
func NotFoundMessage(c *gin.Context, message ...string) {
	msg := MsgResourceNotFound
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusNotFound, msg)
}

// Conflict 409 资源冲突
func Conflict(c *gin.Context, message ...string) {
	msg := MsgResourceConflict
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusConflict, msg)
}

// TooManyRequests 429 请求过多
func TooManyRequests(c *gin.Context) {
	Failure(c, http.StatusTooManyRequests, MsgRateLimitExceeded)
}

// MethodNotAllowed 405 方法不允许
func MethodNotAllowed(c *gin.Context, message ...string) {
	msg := MsgMethodNotAllowed
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusMethodNotAllowed, msg)
}

// NotAcceptable 406 无法满足 Accept 头要求
func NotAcceptable(c *gin.Context, message ...string) {
	msg := MsgNotAcceptable
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusNotAcceptable, msg)
}

// RequestTimeout 408 请求超时
func RequestTimeout(c *gin.Context, message ...string) {
	msg := MsgRequestTimeout
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusRequestTimeout, msg)
}

// Gone 410 资源已永久删除
func Gone(c *gin.Context, message ...string) {
	msg := MsgResourceGone
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusGone, msg)
}

// PayloadTooLarge 413 请求体过大
func PayloadTooLarge(c *gin.Context, message ...string) {
	msg := MsgPayloadTooLarge
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusRequestEntityTooLarge, msg)
}

// UnsupportedMediaType 415 不支持的媒体类型
func UnsupportedMediaType(c *gin.Context, message ...string) {
	msg := MsgUnsupportedMediaType
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusUnsupportedMediaType, msg)
}

// UnprocessableEntity 422 无法处理的实体（业务逻辑验证失败）
func UnprocessableEntity(c *gin.Context, details any, message ...string) {
	msg := MsgUnprocessableEntity
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusUnprocessableEntity, msg, details)
}

// PreconditionFailed 412 预处理失败
func PreconditionFailed(c *gin.Context, message ...string) {
	msg := MsgPreconditionFailed
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Failure(c, http.StatusPreconditionFailed, msg)
}

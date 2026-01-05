package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ============================================================================
// 成功响应函数 (2xx)
// ============================================================================

// OK 200 成功响应
func OK(c *gin.Context, data any, message ...string) {
	msg := MsgSuccess
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Success(c, http.StatusOK, msg, data)
}

// Created 201 创建成功响应
func Created(c *gin.Context, data any, message ...string) {
	msg := MsgCreated
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Success(c, http.StatusCreated, msg, data)
}

// Accepted 202 请求已接受（异步处理中）
func Accepted(c *gin.Context, data any, message ...string) {
	msg := MsgAccepted
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Success(c, http.StatusAccepted, msg, data)
}

// NoContent 204 无内容响应
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// ResetContent 205 重置内容（提示客户端重置视图）
func ResetContent(c *gin.Context) {
	c.Status(http.StatusResetContent)
}

// PartialContent 206 部分内容（范围请求）
func PartialContent(c *gin.Context, data any, message ...string) {
	msg := MsgPartialContent
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	Success(c, http.StatusPartialContent, msg, data)
}

// List 200 列表响应（带分页）
func List(c *gin.Context, data any, meta *PaginationMeta, message ...string) {
	msg := MsgSuccess
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	c.JSON(http.StatusOK, ListResponse{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
		Meta:    meta,
	})
}

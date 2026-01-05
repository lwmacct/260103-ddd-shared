package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ============================================================================
// 底层响应函数
// ============================================================================

// Success 统一成功响应
// 返回格式：{ code: 200, message: "...", data: {...} }
func Success(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, UnifiedResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}

// Failure 统一错误响应
// 返回格式：{ code: 400, message: "...", error: {...} }
func Failure(c *gin.Context, statusCode int, message string, errorDetails ...any) {
	resp := UnifiedResponse{
		Code:    statusCode,
		Message: message,
	}

	if len(errorDetails) > 0 {
		resp.Error = errorDetails[0]
	}

	c.JSON(statusCode, resp)
}

// ============================================================================
// 工具函数
// ============================================================================

// NewPaginationMeta 创建分页元数据
func NewPaginationMeta(total, page, perPage int) *PaginationMeta {
	totalPages := max((total+perPage-1)/perPage, 1)

	meta := &PaginationMeta{
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
		HasMore:    page < totalPages,
	}

	if page > totalPages {
		meta.Warning = fmt.Sprintf("page %d exceeds total_pages %d", page, totalPages)
	}

	return meta
}

package response

// ============================================================================
// 响应结构定义
// ============================================================================

// UnifiedResponse 统一响应结构
// 前端期望格式：{ code: number, message: string, data?: any, error?: any }
type UnifiedResponse struct {
	Code    int    `json:"code"`            // HTTP 状态码（数字）
	Message string `json:"message"`         // 消息描述
	Data    any    `json:"data,omitempty"`  // 成功时的数据
	Error   any    `json:"error,omitempty"` // 失败时的错误详情
}

// Response 通用响应结构（用于 Swagger 文档 - 已废弃，请使用泛型版本）
//
// Deprecated: 使用 DataResponse[T] 替代
type Response struct {
	Message string `json:"message,omitempty"` // 消息
	Data    any    `json:"data,omitempty"`    // 数据
}

// DataResponse 泛型数据响应（用于 Swagger 文档，消除 allOf 嵌套）
//
//	@Description	统一数据响应格式
type DataResponse[T any] struct {
	Code    int    `json:"code"`            // HTTP 状态码
	Message string `json:"message"`         // 消息描述
	Data    T      `json:"data,omitempty"`  // 响应数据
	Error   any    `json:"error,omitempty"` // 错误详情（仅失败时）
}

// MessageResponse 纯消息响应（无数据）
//
//	@Description	纯消息响应格式
type MessageResponse struct {
	Code    int    `json:"code"`    // HTTP 状态码
	Message string `json:"message"` // 消息描述
}

// ErrorResponse 错误响应结构（用于 Swagger 文档）
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail 错误详情
type ErrorDetail struct {
	Code    string `json:"code"`              // 业务错误码（小写下划线）
	Message string `json:"message"`           // 错误消息
	Details any    `json:"details,omitempty"` // 额外详情（如验证错误列表）
}

// ListResponse 列表响应（带分页信息 - 已废弃，请使用泛型版本）
//
// Deprecated: 使用 PagedResponse[T] 替代
type ListResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    any             `json:"data"`
	Meta    *PaginationMeta `json:"meta,omitempty"`
}

// PagedResponse 泛型分页响应（用于 Swagger 文档，消除 allOf 嵌套）
//
//	@Description	分页列表响应格式
type PagedResponse[T any] struct {
	Code    int             `json:"code"`           // HTTP 状态码
	Message string          `json:"message"`        // 消息描述
	Data    []T             `json:"data"`           // 数据列表
	Meta    *PaginationMeta `json:"meta,omitempty"` // 分页信息
}

// PaginationMeta 分页元数据
type PaginationMeta struct {
	Total      int    `json:"total"`                 // 总记录数
	Page       int    `json:"page"`                  // 当前页码
	PerPage    int    `json:"per_page"`              // 每页数量
	TotalPages int    `json:"total_pages,omitempty"` // 总页数
	HasMore    bool   `json:"has_more,omitempty"`    // 是否有下一页
	Warning    string `json:"warning,omitempty"`     // 页码越界警告
}

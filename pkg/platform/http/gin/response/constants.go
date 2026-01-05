package response

const (
	// ============================================================================
	// 成功响应消息 (2xx)
	// ============================================================================

	// MsgSuccess 表示操作成功
	MsgSuccess = "操作成功"

	// MsgCreated 表示资源创建成功
	MsgCreated = "创建成功"

	// MsgUpdated 表示资源更新成功
	MsgUpdated = "更新成功"

	// MsgDeleted 表示资源删除成功
	MsgDeleted = "删除成功"

	// MsgAccepted 表示请求已接受，将在后台处理
	MsgAccepted = "请求已接受"

	// MsgPartialContent 表示部分内容（范围请求）
	MsgPartialContent = "部分内容"

	// ============================================================================
	// 客户端错误消息 (4xx)
	// ============================================================================

	// MsgValidationFailed 表示请求参数验证失败
	MsgValidationFailed = "验证失败"

	// MsgAuthenticationRequired 表示需要身份认证
	MsgAuthenticationRequired = "请先登录"

	// MsgAccessForbidden 表示无权限访问
	MsgAccessForbidden = "无权访问"

	// MsgResourceNotFound 表示请求的资源不存在
	MsgResourceNotFound = "资源不存在"

	// MsgNotFoundFormat 资源未找到消息格式（使用 fmt.Sprintf）
	MsgNotFoundFormat = "%s not found"

	// MsgResourceConflict 表示资源冲突
	MsgResourceConflict = "资源冲突"

	// MsgResourceGone 表示资源已永久删除
	MsgResourceGone = "资源已删除"

	// MsgMethodNotAllowed 表示不允许的 HTTP 方法
	MsgMethodNotAllowed = "方法不支持"

	// MsgNotAcceptable 表示无法满足 Accept 头要求
	MsgNotAcceptable = "无法接受的内容类型"

	// MsgRequestTimeout 表示请求超时
	MsgRequestTimeout = "请求超时"

	// MsgPayloadTooLarge 表示请求体过大
	MsgPayloadTooLarge = "请求体过大"

	// MsgUnsupportedMediaType 表示不支持的媒体类型
	MsgUnsupportedMediaType = "不支持的媒体类型"

	// MsgUnprocessableEntity 表示请求格式正确但语义错误
	MsgUnprocessableEntity = "无法处理请求"

	// MsgPreconditionFailed 表示预处理失败（如条件请求）
	MsgPreconditionFailed = "预处理失败"

	// MsgRateLimitExceeded 表示请求过于频繁
	MsgRateLimitExceeded = "请求过于频繁"

	// ============================================================================
	// 服务端错误消息 (5xx)
	// ============================================================================

	// MsgInternalError 表示服务器内部错误
	MsgInternalError = "服务器内部错误"

	// MsgNotImplemented 表示功能未实现
	MsgNotImplemented = "功能未实现"

	// MsgServiceUnavailable 表示服务暂时不可用
	MsgServiceUnavailable = "服务暂时不可用"
)

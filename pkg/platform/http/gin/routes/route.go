package routes

import "github.com/gin-gonic/gin"

// Method HTTP 方法类型
type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
)

// Route 路由定义（声明式）
type Route struct {
	Handlers []gin.HandlerFunc `json:""` // 处理链：[0] 处理函数，后续为中间件（不可序列化）

	// 基本
	Audit       bool   `json:"audit"`       // 是否开启审计
	Method      Method `json:"method"`      // HTTP 方法
	Path        string `json:"path"`        // Gin 路由路径
	OperationID string `json:"operationId"` // OpenAPI operationId + 权限标识

	// 文档
	Tags        []string `json:"tags"`        // 分类标签（Swagger tags）
	Summary     string   `json:"summary"`     // 简短标题（Swagger summary）
	Description string   `json:"description"` // 详细说明（Swagger description，支持 CommonMark）
	Deprecated  bool     `json:"deprecated"`  // 是否废弃（Swagger UI 会特殊标记）
}

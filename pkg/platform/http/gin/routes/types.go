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
	// 基本信息
	Method      Method            // HTTP 方法
	Path        string            // Gin 路由路径
	Operation   string            // Operation: domain:resource:action
	Handler     gin.HandlerFunc   // 处理函数
	Middlewares []gin.HandlerFunc // 中间件列表

	// Swagger 文档
	Tags        string
	Summary     string
	Description string
}

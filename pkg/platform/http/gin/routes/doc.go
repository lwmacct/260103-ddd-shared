// Package routes 提供声明式路由定义和路径转换工具。
//
// # Overview
//
// 本包定义了与框架无关的路由元数据结构，采用 OpenAPI 规范格式：
//   - [Route]: 路由定义（路径、方法、处理函数、文档元数据）
//   - [Method]: HTTP 方法类型常量
//   - [ToGinPath]: OpenAPI 路径转 Gin 格式
//   - [ToOpenAPIPath]: Gin 路径转 OpenAPI 格式
//
// # 路径格式
//
// 路由路径采用 OpenAPI 规范格式，参数使用花括号：
//
//	/users/{id}                    // 单个参数
//	/users/{userId}/posts/{postId} // 多个参数
//
// 注册到 Gin 时需转换为冒号格式：
//
//	ginPath := routes.ToGinPath("/users/{id}") // => "/users/:id"
//
// # Usage
//
//	route := routes.Route{
//	    Method:      routes.GET,
//	    Path:        "/users/{id}",
//	    OperationID: "iam:user:get",
//	    Summary:     "获取用户详情",
//	}
//
//	// 公开接口（无需认证）
//	loginRoute := routes.Route{
//	    Method:      routes.POST,
//	    Path:        "/auth/login",
//	    OperationID: "iam:auth:login",
//	    Public:      true,
//	}
//
//	// 文件上传接口
//	uploadRoute := routes.Route{
//	    Method:   routes.POST,
//	    Path:     "/files",
//	    Consumes: []string{"multipart/form-data"},
//	}
//
//	// 限流接口
//	smsRoute := routes.Route{
//	    Method:    routes.POST,
//	    Path:      "/sms/send",
//	    RateLimit: "10/m", // 每分钟 10 次
//	}
//
//	// 注册到 Gin
//	r.Handle(string(route.Method), routes.ToGinPath(route.Path), route.Handlers...)
package routes

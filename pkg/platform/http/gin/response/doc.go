// Package response 提供统一的 HTTP 响应格式。
//
// 本包定义了前后端约定的响应结构，确保 API 响应的一致性：
//
// 成功响应格式：
//
//	{
//	  "code": 200,
//	  "message": "操作成功",
//	  "data": { ... }
//	}
//
// 错误响应格式：
//
//	{
//	  "code": 400,
//	  "message": "验证失败",
//	  "error": { ... }
//	}
//
// 列表响应格式：
//
//	{
//	  "code": 200,
//	  "message": "操作成功",
//	  "data": [...],
//	  "meta": { "total": 100, "page": 1, "per_page": 10 }
//	}
//
// 使用示例：
//
//	response.OK(c, user)                           // 使用默认消息 "操作成功"
//	response.OK(c, user, "登录成功")                // 自定义消息
//	response.Created(c, user)                      // 使用默认消息 "创建成功"
//	response.BadRequest(c, "无效输入")              // 错误响应
//	response.UnprocessableEntity(c, details)       // 业务验证失败（带详情）
//	response.UnprocessableEntity(c, details, "库存不足") // 业务验证失败（自定义消息）
//	response.List(c, users, meta)                  // 列表响应（使用默认消息）
//	response.List(c, users, meta, "查询成功")      // 列表响应（自定义消息）
package response

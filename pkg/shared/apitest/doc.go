// Package apitest 提供 HTTP API 集成测试工具。
//
// 本包提供了类型安全的 HTTP 测试客户端，以及常用的测试辅助函数。
//
// # 基本用法
//
//	client := apitest.NewClient("http://localhost:8080")
//
//	// 设置认证令牌
//	client.SetToken("your-token-here")
//
//	// GET 请求
//	user, err := apitest.Get[UserDTO](client, "/api/users/1", nil)
//
//	// GET 分页列表请求
//	users, meta, err := apitest.GetList[UserDTO](client, "/api/users", nil)
//
//	// POST 请求
//	result, err := apitest.Post[CreateResultDTO](client, "/api/users", createReq)
//
//	// PUT 请求
//	updated, err := apitest.Put[UserDTO](client, "/api/users/1", updateReq)
//
//	// PATCH 请求
//	patched, err := apitest.Patch[UserDTO](client, "/api/users/1", patchReq)
//
//	// DELETE 请求（无返回值）
//	err = client.Delete("/api/users/1")
//
//	// DELETE 请求（有返回值，如软删除）
//	deleted, err := apitest.DeleteOne[UserDTO](client, "/api/users/1", nil)
//
// # 条件跳过测试
//
// 使用 SkipIfNotAPITest 在未设置 API_TEST 环境变量时跳过测试：
//
//	func TestSomething(t *testing.T) {
//	    apitest.SkipIfNotAPITest(t)
//	    // ... 需要运行服务的测试代码
//	}
//
// # 数据提取
//
// 使用泛型提取函数从切片中安全地提取字段：
//
//	ids := apitest.ExtractIDs(users, func(u UserDTO) uint { return u.ID })
//	names := apitest.ExtractStrings(users, func(u UserDTO) string { return u.Username })
//
// # BC 特定扩展
//
// 各个限界上下文可以嵌入 Client 结构体以添加 BC 特定方法（如认证辅助函数）：
//
//	type IAMClient struct {
//	    apitest.Client
//	    devSecret string
//	}
//
//	// 添加 BC 特定方法
//	func (c *IAMClient) Login(account, password string) (*LoginResponseDTO, error) {
//	    // ... 实现登录逻辑
//	    return resp, nil
//	}
//
// # 使用方式
//
//	// 仅在设置了 API_TEST 环境变量时运行测试
//	API_TEST=1 go test -v -count=1 ./internal/manualtest/...
package apitest

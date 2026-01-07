# Go DDD Platform Library

基于领域驱动设计（DDD）和 CQRS 模式的 Go 平台基础设施库，提供构建企业级应用所需的核心组件。

## 特性

- **HTTP/Gin 集成**：路由元数据、标准化响应、中间件（CORS、日志、RequestID）
- **RedisJSON 缓存**：类型安全的 JSON 存储、Pipeline 操作
- **数据库管理**：GORM 连接池、自动迁移、种子数据
- **事件总线**：同步发布/订阅、通配符匹配
- **任务队列**：Redis FIFO 队列、多 Worker 并行
- **RBAC 权限**：URN 风格操作标识、权限解析器
- **可观测性**：OpenTelemetry 追踪

## 技术栈

| 组件     | 技术                  |
| -------- | --------------------- |
| Web 框架 | Gin 1.11              |
| ORM      | GORM 1.31             |
| 缓存     | Redis 8.x + RedisJSON |
| 依赖注入 | Uber Fx               |
| 追踪     | OpenTelemetry         |

## 目录结构

```
pkg/
├── platform/                    # 平台基础设施层
│   ├── cache/                   # Redis 缓存服务
│   ├── db/                      # 数据库连接与迁移
│   ├── eventbus/                # 事件总线
│   ├── health/                  # 健康检查
│   ├── http/gin/                # Gin HTTP 集成
│   │   ├── server.go            # HTTP 服务器包装器
│   │   ├── ctxutil/             # Context 工具
│   │   ├── handler/             # 通用 Handler 模式
│   │   ├── helper/              # HTTP 辅助工具
│   │   ├── middleware/          # CORS, 日志, RequestID
│   │   ├── permission/          # RBAC 权限解析
│   │   ├── response/            # 标准化响应
│   │   └── routes/              # 路由元数据
│   ├── queue/                   # Redis 队列
│   └── telemetry/               # OpenTelemetry
│
└── shared/                      # 共享组件
    ├── cache/                   # 缓存接口与错误
    ├── captcha/                 # 验证码服务
    ├── event/                   # 事件类型定义
    └── health/                  # 健康检查接口
```

## 核心组件

### Response System

标准化 HTTP 响应格式：

```go
import "pkg/platform/http/gin/response"

// 单对象响应
response.DataResponse[UserDTO]{Data: user}

// 分页列表响应
response.PagedResponse[UserDTO]{
    Items: users,
    Total: 100,
    Page:  1,
    Size:  10,
}

// 错误响应
response.ErrorResponse{Code: 400, Message: "参数错误"}

// 204 无内容响应（用于 Swagger 文档）
response.EmptyResponse{}
```

### Routes

框架无关的路由元数据：

```go
import "pkg/platform/http/gin/routes"

route := routes.Route{
    Method:      routes.GET,
    Path:        "/api/users/{id}",  // OpenAPI 风格
    OperationID: "iam:user:read",
    Summary:     "获取用户详情",
    Tags:        []string{"user"},
    Public:      false,
}

// 路径转换
ginPath := routes.ToGinPath(route.Path)      // "/api/users/:id"
openAPIPath := routes.ToOpenAPIPath(ginPath) // "/api/users/{id}"
```

### Cache

RedisJSON 缓存服务：

```go
import "pkg/platform/cache"

// 创建缓存服务
svc := cache.NewUserCacheService(redisClient, "app:")

// 写入（Pipeline：JSON.SET + EXPIRE）
svc.Set(ctx, userID, userData, 24*time.Hour)

// 读取（自动解包 JSON 数组）
user, err := svc.Get(ctx, userID)
```

### EventBus

同步事件发布/订阅：

```go
import "pkg/platform/eventbus"

bus := eventbus.New()

// 订阅（支持通配符）
bus.Subscribe("user.*", handler)
bus.Subscribe("*", auditLogger)

// 发布
bus.Publish(ctx, "user.created", event)
```

### Queue

Redis 任务队列：

```go
import "pkg/platform/queue"

q := queue.NewRedisQueue(redisClient, "tasks")

// 生产者
q.Push(ctx, taskData)

// 消费者（阻塞）
q.Process(ctx, func(data []byte) error {
    return handleTask(data)
})
```

### Permission

RBAC 权限解析：

```go
import "pkg/platform/http/gin/permission"

resolver := permission.NewResolver(permissionService)

// 检查权限
hasAccess := resolver.Check(ctx, userID, "iam:user:write")

// URN 格式：{scope}:{resource}:{action}
// 示例：iam:user:read, app:setting:write
```

## 使用方式

### 安装

```bash
go get github.com/your-org/ddd-platform
```

### 基本集成

```go
package main

import (
    "pkg/platform/cache"
    "pkg/platform/db"
    "pkg/platform/http/gin/response"
    "pkg/platform/http/gin/routes"
)

func main() {
    // 1. 初始化数据库
    database := db.NewConnection(cfg.Database)

    // 2. 初始化缓存
    redisClient := cache.NewRedisClient(cfg.Redis)

    // 3. 定义路由
    userRoutes := []routes.Route{
        {Method: routes.GET, Path: "/api/users", OperationID: "iam:user:list"},
        {Method: routes.POST, Path: "/api/users", OperationID: "iam:user:create"},
    }

    // 4. 使用响应工具
    func GetUser(c *gin.Context) {
        user := fetchUser(c.Param("id"))
        c.JSON(200, response.DataResponse[UserDTO]{Data: user})
    }
}
```

### Uber Fx 集成

```go
import "go.uber.org/fx"

func Module() fx.Option {
    return fx.Module("platform",
        fx.Provide(
            db.NewConnection,
            cache.NewRedisClient,
            eventbus.New,
        ),
    )
}
```

## 开发命令

```bash
# 启动依赖服务
docker compose -p srv-db up -d

# 单元测试
go test ./...

# 编译检查
go build -o /dev/null ./...

# Lint 检查
golangci-lint run --new
```

## License

MIT

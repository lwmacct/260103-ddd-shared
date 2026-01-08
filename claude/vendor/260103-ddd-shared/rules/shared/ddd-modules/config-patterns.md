# 模块配置模式规范

## 配置分层原则

### 1. 配置分类

**Platform 配置**（基础设施，跨模块共享）：
- Redis、DB、EventBus、Queue、Telemetry
- 位置：`internal/config` 的 Data/Server/Telemetry 等
- 注入：直接注入基础设施对象（`*redis.Client`, `*gorm.DB`）

**Module 配置**（业务模块特定）：
- HTTP 路由、业务参数、功能开关
- 位置：`internal/config.Modules.{BC}` + `pkg/modules/{bc}/config`
- 注入：通过容器层转换为模块 Config

### 2. 为什么需要区分？

| 维度 | Platform 配置 | Module 配置 |
|------|--------------|-------------|
| **作用域** | 跨所有 BC 共享 | 单个 BC 特定 |
| **管理责任** | 容器层（internal/container） | 业务模块 |
| **示例** | Redis 连接、DB URL | HTTP BasePath、业务参数 |
| **注入方式** | 直接注入对象 | 通过 Config 结构 |

**关键原则**：
- ✅ Platform 配置由容器层统一管理
- ✅ Module 配置定义在模块内部，但值来自全局配置
- ❌ 业务模块不应配置 Platform 层参数（如 Redis 前缀）

## 模块 Config 包规范

### 职责

**应该做的**：
- ✅ 定义模块配置结构（作为配置契约）
- ✅ 包文档说明配置用途
- ✅ 使用 `koanf` tags 支持配置映射

**不应该做的**：
- ❌ 定义默认值（默认值在 `internal/config.DefaultConfig()`）
- ❌ 包含 Platform 层配置（Redis、DB、Telemetry 等）
- ❌ 依赖 `internal/config`（保持依赖倒置）

### 目录结构

```
pkg/modules/{bc}/config/
├── doc.go        # 包文档
└── config.go     # 配置结构定义
```

### 示例

**正确示例**：

```go
// Package config 提供 Settings 模块的配置定义。
//
// 本包定义 Settings 模块所需的所有配置项，实现模块配置自治。
// 配置通过依赖注入由 internal/container 提供，Settings 模块不直接依赖 internal/config。
package config

// Config Settings 模块配置
type Config struct {
    // BasePath Settings API 的基础路径
    BasePath string `koanf:"base-path"`
}
```

**错误示例**：

```go
// ❌ 错误：包含 Platform 层配置
type Config struct {
    Redis struct {
        KeyPrefix string  // 这是 Platform 配置，不应在这里
    }
    BasePath string
}

// ❌ 错误：定义默认值
func DefaultConfig() Config {
    return Config{
        BasePath: "/api/admin/settings",  // 默认值应在 internal/config
    }
}
```

## 容器层配置转换规范

### 模式

每个需要配置的模块在 `internal/container` 提供转换模块：

```go
// internal/container/settings.go
package container

import (
    "go.uber.org/fx"

    "github.com/lwmacct/260103-ddd-bc-settings/internal/config"
    settingsconfig "github.com/lwmacct/260103-ddd-bc-settings/pkg/modules/settings/config"
)

// SettingsConfigModule 提供 Settings 模块的配置。
var SettingsConfigModule = fx.Module("settings.config",
    fx.Provide(newSettingsConfig),
)

// newSettingsConfig 从全局配置构造 Settings 模块配置。
func newSettingsConfig(cfg *config.Config) *settingsconfig.Config {
    return &settingsconfig.Config{
        BasePath: cfg.Settings.BasePath,
    }
}
```

### 关键点

1. **模块命名**：`{BC}ConfigModule`
2. **函数命名**：`new{BC}Config`
3. **参数类型**：`*config.Config`（全局配置）
4. **返回类型**：`*{bc}config.Config`（模块配置）

## Platform 层配置直接注入

### 原则

Platform 层基础设施对象应该直接注入到需要的地方，不通过模块 Config 包装。

### 示例

**正确示例**：

```go
// ✅ 正确：直接注入 Platform 层对象
func NewSettingsCacheServiceAs(
    client *redis.Client,     // Platform 层对象
    keyPrefix string,         // Platform 层配置
) CacheResult {
    service := &settingsCacheService{
        client:    client,
        keyPrefix: keyPrefix,
    }
    return CacheResult{...}
}
```

```go
// ✅ 正确：容器层直接提供 Platform 配置
func provideRedis(cfg *config.Config) (*redis.Client, string, error) {
    client := redis.NewClient(&redis.Options{...})
    return client, cfg.Data.RedisKeyPrefix, nil
}
```

**错误示例**：

```go
// ❌ 错误：通过模块 Config 传递 Platform 配置
func NewSettingsCacheServiceAs(
    settingsCfg *settingsconfig.Config,  // 包含 Redis 配置
) CacheResult {
    client := redis.NewClient(...)  // 不应在模块中创建客户端
    return CacheResult{...}
}
```

```go
// ❌ 错误：模块 Config 包含 Platform 字段
type Config struct {
    RedisKeyPrefix string  // 这是 Platform 配置
    BasePath       string
}
```

## 全局配置组织

### 结构

```go
// internal/config/config.go
type Config struct {
    // Platform 配置（跨模块共享）
    Server    Server
    Data      Data      // Redis, DB
    Telemetry Telemetry

    // Module 配置（按 BC 分组）
    Settings  Settings  // Settings 模块配置
    IAM       IAM       // IAM 模块配置
    // ...
}
```

### 默认值

所有默认值统一在 `internal/config.DefaultConfig()` 中定义：

```go
func DefaultConfig() Config {
    return Config{
        Data: Data{
            RedisURL:       "redis://localhost:6379/0",
            RedisKeyPrefix: "app:",  // Platform 默认值
        },
        Settings: Settings{
            BasePath: "/api/admin/settings",  // Module 默认值
        },
    }
}
```

## 常见错误

| 错误 | 症状 | 解决方案 |
|------|------|----------|
| 模块 Config 包含 Redis 字段 | 配置冗余，难以维护 | 移除到容器层，直接注入 |
| 模块 Config 定义默认值 | 配置分散，不符合单一来源原则 | 在 internal/config 定义 |
| 容器层传递整个模块 Config | 包含不必要的 Platform 配置 | 只传递业务相关字段 |
| 模块直接依赖 internal/config | 破坏依赖倒置 | 通过容器层转换 |

## 检查清单

### 模块 Config 包
- [ ] 定义在 `pkg/modules/{bc}/config/`
- [ ] `doc.go` 包文档存在
- [ ] 不定义默认值
- [ ] 不包含 Platform 层配置（Redis、DB 等）
- [ ] 使用 `koanf` tags 支持配置映射

### 容器层
- [ ] 提供 `{BC}ConfigModule`
- [ ] 提供 `new{BC}Config()` 转换函数
- [ ] Platform 层对象直接注入（不通过模块 Config）

### 全局配置
- [ ] 所有默认值在 `internal/config.DefaultConfig()`
- [ ] Platform 配置和 Module 配置清晰分组
- [ ] 配置结构单一来源

## 示例项目结构

```
internal/
├── config/
│   └── config.go          # 全局配置（Platform + Module）
└── container/
    ├── infra.go           # Platform 层提供（Redis, DB）
    ├── settings.go        # Settings 配置转换
    └── iam.go             # IAM 配置转换

pkg/modules/settings/
├── config/
│   ├── doc.go             # 包文档
│   └── config.go          # 只有 BasePath 等业务配置
└── infra/cache/
    └── module.go          # 直接注入 redis.Client + keyPrefix
```

## 迁移指南

如果现有项目存在配置混用问题，按以下步骤重构：

1. **识别 Platform 配置**：找出 Redis、DB 等 Platform 字段
2. **移动到全局配置**：将这些字段移到 `internal/config`
3. **简化模块 Config**：只保留业务相关字段
4. **更新容器层**：修改转换函数，只传递业务配置
5. **更新注入点**：Platform 对象直接注入
6. **删除默认值**：模块 Config 的 `DefaultConfig()` 函数

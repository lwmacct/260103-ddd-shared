package ctxutil

import (
	"github.com/gin-gonic/gin"
)

const (
	Permissions = "permissions"
	OperationID = "operation"
	Roles       = "roles"
	Username    = "username"
	UserID      = "user_id"
	UserRole    = "user_role"
	OrgID       = "org_id"
	TeamID      = "team_id"
	RequestID   = "request_id"
	Email       = "email"
	AuthType    = "auth_type"
	IsAdmin     = "is_admin"
	Locale      = "locale"
	Timezone    = "timezone"
)

// Get 从 Context 安全获取指定 key 的值并断言为类型 T。
// 返回值和成功标志，类似 map 查找的 comma-ok 模式。
func Get[T any](c *gin.Context, key string) (T, bool) {
	var zero T
	value, exists := c.Get(key)
	if !exists {
		return zero, false
	}
	result, ok := value.(T)
	return result, ok
}

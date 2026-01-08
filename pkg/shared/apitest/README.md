# Manualtest Factory Pattern Guide

This guide documents the factory pattern used in manualtest for creating test resources.

## Overview

Each bounded context (BC) defines its own factory functions in `pkg/modules/{bc}/adapters/gin/manualtest/factory.go`. These functions create test resources (users, roles, settings, etc.) and automatically register cleanup handlers.

## Pattern

### Basic Factory Function

Creates a resource and automatically cleans up when the test finishes:

```go
func CreateTestUser(t *testing.T, c *Client, prefix string) *user.UserDTO {
    t.Helper()

    username := fmt.Sprintf("%s_%s", prefix, uuid.New().String()[:8])
    req := user.CreateDTO{
        Username:  username,
        Email:     username + "@test.local",
        Password:  "test123456",
        RealName:  "测试用户",
    }

    resp, err := Post[user.UserDTO](c, "/api/admin/users", req)
    require.NoError(t, err, "创建测试用户失败: %s", username)

    userID := resp.ID
    t.Cleanup(func() {
        if userID > 0 {
            _ = c.Delete(fmt.Sprintf("/api/admin/users/%d", userID))
        }
    })

    return resp
}
```

### With Cleanup Control

Use when the test itself deletes the resource:

```go
func CreateTestUserWithCleanupControl(t *testing.T, c *Client, prefix string) (*user.UserDTO, func()) {
    t.Helper()

    username := fmt.Sprintf("%s_%s", prefix, uuid.New().String()[:8])
    req := user.CreateDTO{...}

    result, err := Post[user.UserDTO](c, "/api/admin/users", req)
    require.NoError(t, err, "创建测试用户失败: %s", username)

    userID := result.ID
    deleted := false

    t.Cleanup(func() {
        if !deleted && userID > 0 {
            _ = c.Delete(fmt.Sprintf("/api/admin/users/%d", userID))
        }
    })

    return result, func() { deleted = true }
}
```

Usage:

```go
func TestDeleteUser(t *testing.T) {
    c := manualtest.LoginAsAdmin(t)

    user, markDeleted := manualtest.CreateTestUserWithCleanupControl(t, c, "del")

    // Test delete logic
    err := c.Delete("/api/admin/users/" + strconv.Itoa(int(user.ID)))
    require.NoError(t, err)

    // Mark as deleted so Cleanup doesn't try again
    markDeleted()
}
```

## Key Points

1. **Use `t.Helper()`** - Marks function as test helper for better error messages
2. **Unique names** - Use `prefix + UUID` or `prefix + timestamp` for uniqueness
3. **Auto cleanup** - Always register `t.Cleanup()` immediately after creation
4. **ID tracking** - Save ID before cleanup to prevent cleanup after deletion
5. **Error handling** - Use `require.NoError()` to fail fast on creation errors

## BC-Specific Variations

Each BC adapts this pattern to its domain:

- **IAM**: `CreateTestUser()`, `CreateTestRole()`
- **Settings**: `CreateTestSetting()`, `CreateTestSettingCategory()`
- **LMS**: `CreateTestCustomer()`, `CreateTestAsset()`, `CreateTestOrder()`

## Naming Convention

- Basic: `CreateTest{Resource}(t, c, prefix) *ResourceDTO`
- With control: `CreateTest{Resource}WithCleanupControl(t, c, prefix) (*ResourceDTO, func())`

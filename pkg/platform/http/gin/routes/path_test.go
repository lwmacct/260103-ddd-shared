package routes_test

import (
	"testing"

	"github.com/lwmacct/260103-ddd-shared/pkg/platform/http/gin/routes"

	"github.com/stretchr/testify/assert"
)

func TestToGinPath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"单个参数", "/users/{id}", "/users/:id"},
		{"多个参数", "/users/{userId}/posts/{postId}", "/users/:userId/posts/:postId"},
		{"无参数", "/users", "/users"},
		{"空字符串", "", ""},
		{"只有参数", "{id}", ":id"},
		{"连续参数", "/{a}/{b}/{c}", "/:a/:b/:c"},
		{"复杂路径", "/api/v1/users/{userId}/settings/{key}", "/api/v1/users/:userId/settings/:key"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := routes.ToGinPath(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToOpenAPIPath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"单个参数", "/users/:id", "/users/{id}"},
		{"多个参数", "/users/:userId/posts/:postId", "/users/{userId}/posts/{postId}"},
		{"无参数", "/users", "/users"},
		{"空字符串", "", ""},
		{"只有参数", ":id", "{id}"},
		{"连续参数", "/:a/:b/:c", "/{a}/{b}/{c}"},
		{"复杂路径", "/api/v1/users/:userId/settings/:key", "/api/v1/users/{userId}/settings/{key}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := routes.ToOpenAPIPath(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRoundTrip(t *testing.T) {
	// OpenAPI -> Gin -> OpenAPI
	openAPIPath := "/users/{userId}/posts/{postId}"
	ginPath := routes.ToGinPath(openAPIPath)
	backToOpenAPI := routes.ToOpenAPIPath(ginPath)
	assert.Equal(t, openAPIPath, backToOpenAPI)

	// Gin -> OpenAPI -> Gin
	ginPath2 := "/users/:userId/posts/:postId"
	openAPIPath2 := routes.ToOpenAPIPath(ginPath2)
	backToGin := routes.ToGinPath(openAPIPath2)
	assert.Equal(t, ginPath2, backToGin)
}

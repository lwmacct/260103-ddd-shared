package routes

import "regexp"

// OpenAPIParamRegex 匹配 OpenAPI 风格路径参数 {param}。
var OpenAPIParamRegex = regexp.MustCompile(`\{([^}]+)\}`)

// GinParamRegex 匹配 Gin 风格路径参数 :param。
var GinParamRegex = regexp.MustCompile(`:([^/]+)`)

// ToGinPath 将 OpenAPI 风格路径参数 {param} 转换为 Gin 风格 :param。
//
//	ToGinPath("/users/{id}") => "/users/:id"
//	ToGinPath("/users/{userId}/posts/{postId}") => "/users/:userId/posts/:postId"
func ToGinPath(path string) string {
	return OpenAPIParamRegex.ReplaceAllString(path, ":$1")
}

// ToOpenAPIPath 将 Gin 风格路径参数 :param 转换为 OpenAPI 风格 {param}。
//
//	ToOpenAPIPath("/users/:id") => "/users/{id}"
//	ToOpenAPIPath("/users/:userId/posts/:postId") => "/users/{userId}/posts/{postId}"
func ToOpenAPIPath(path string) string {
	return GinParamRegex.ReplaceAllString(path, "{$1}")
}

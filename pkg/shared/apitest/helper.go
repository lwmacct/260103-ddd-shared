package apitest

import (
	"os"
	"testing"
)

// SkipIfNotAPITest 如果 API_TEST 环境变量未设置则跳过测试。
func SkipIfNotAPITest(t *testing.T) {
	t.Helper()
	if os.Getenv("API_TEST") == "" {
		t.SkipNow()
	}
}

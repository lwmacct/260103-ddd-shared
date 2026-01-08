package infrastructure

import (
	"fmt"
	"image/color"
	"time"

	"github.com/mojocn/base64Captcha"

	"github.com/lwmacct/260103-ddd-shared/pkg/shared/captcha"
)

// 验证码默认配置
const (
	DefaultExpiration = 5 * time.Minute
	DefaultLength     = 6   // 验证码长度
	DefaultWidth      = 120 // 图片宽度
	DefaultHeight     = 40  // 图片高度
)

// 数字验证码字符集
const charset = "0123456789"

// Service 验证码生成服务实现
type Service struct {
	driver *base64Captcha.DriverString
}

// NewService 创建验证码服务
func NewService() captcha.Service {
	// 创建字符串验证码驱动（支持数字）
	driver := base64Captcha.NewDriverString(
		DefaultHeight,                     // 高度
		DefaultWidth,                      // 宽度
		80,                                // 干扰噪点数量
		base64Captcha.OptionShowSlimeLine, // 显示干扰线
		DefaultLength,                     // 验证码长度
		charset,                           // 字符集（纯数字）
		&color.RGBA{245, 245, 245, 255},   // 背景颜色（浅灰）
		nil,                               // 使用默认字体存储
		[]string{"wqy-microhei.ttc"},      // 字体列表
	).ConvertFonts()

	return &Service{
		driver: driver,
	}
}

// GenerateRandomCode 生成随机验证码
// 返回 (captchaID, imageBase64, code, error)
func (s *Service) GenerateRandomCode() (string, string, string, error) {
	captchaInstance := base64Captcha.NewCaptcha(s.driver, base64Captcha.DefaultMemStore)
	captchaID, b64s, _, err := captchaInstance.Generate()
	if err != nil {
		return "", "", "", fmt.Errorf("failed to generate captcha: %w", err)
	}

	// 从 base64Captcha 的 store 获取验证码值
	code := base64Captcha.DefaultMemStore.Get(captchaID, false)

	return captchaID, b64s, code, nil
}

// GenerateCustomCodeImage 生成指定文本的验证码图片
// 用于开发模式
func (s *Service) GenerateCustomCodeImage(text string) (string, error) {
	item, err := s.driver.DrawCaptcha(text)
	if err != nil {
		return "", fmt.Errorf("failed to draw captcha: %w", err)
	}
	return item.EncodeB64string(), nil
}

// GenerateDevCaptchaID 生成开发模式验证码ID
func (s *Service) GenerateDevCaptchaID() string {
	return fmt.Sprintf("dev-%d", time.Now().Unix())
}

// GetDefaultExpiration 获取默认过期时间（秒）
func (s *Service) GetDefaultExpiration() int64 {
	return int64(DefaultExpiration.Seconds())
}

// ============================================================================
// 确保实现接口
// ============================================================================

var _ captcha.Service = (*Service)(nil)

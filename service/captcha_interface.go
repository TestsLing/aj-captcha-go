package service

type CaptchaInterface interface {

	// Get 获取验证码
	Get() map[string]interface{}

	// Check 核对验证码
	Check(token string, pointJson string) error

	// Verification 二次校验验证码(后端)
	Verification(token string, pointJson string) error
}

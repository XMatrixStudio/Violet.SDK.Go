package violet

import "errors"

// Error 错误列表
var (
	ErrorInvalidKey    = errors.New("invalid_key")    // Key 无效
	ErrorUnableDecrypt = errors.New("unable_decrypt") // 无法解密
	ErrorTimeoutCode   = errors.New("timeout_code")   // Code 已超时
	ErrorInvalidCode   = errors.New("invalid_code")   // Code 无效
	ErrorTimeoutToken  = errors.New("timeout_token")  // Token 已超时
	ErrorInvalidToken  = errors.New("invalid_token")  // Token 无效
	ErrorTimeoutState  = errors.New("timeout_state")  // State 已超时
	ErrorInvalidState  = errors.New("invalid_state")  // State 无效
	ErrorNetwork       = errors.New("network_error")  // 网络错误
	ErrorServer        = errors.New("server_error")   // 服务器错误
	ErrorUnknown       = errors.New("unknown")        // 未知错误
)

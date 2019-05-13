package violet

// TokenRes Token结果
type TokenRes struct {
	UserID string
	Token  string
}

// ErrorRes 错误返回值
type ErrorRes struct {
	Error string
}

// UserInfoRes 用户信息
type UserInfoRes struct {
	ID       string
	Avatar   string
	Bio      string
	Birthday string
	Email    string
	Location string
	Nickname string
	Phone    string
	URL      string
}

// LoginRes 登陆返回值
type LoginRes struct {
	Valid bool
	Email string
	Code  string
}

// ScopeType 请求权限
type ScopeType string

// ScopeType 请求权限种类
const (
	scopeBase  ScopeType = "base"
	ScopeInfo  ScopeType = "info"
	ScopeEmail ScopeType = "email"
)

func (s *ScopeTypes) String() (scopes []string) {
	for _, t := range *s {
		scopes = append(scopes, string(t))
	}
	return
}

type ScopeTypes []ScopeType

// AuthOption 授权选项
type AuthOption struct {
	Scopes    ScopeTypes // 请求额外权限列表
	QuickMode bool       // 快速授权模式， 默认开启
}
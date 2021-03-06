package violetSdk

// Config 配置文件
type Config struct {
	ClientID string `yaml:"ClientID"`
	// ClientKey KEY
	ClientKey string `yaml:"ClientKey"`
	// ServerHost 服务器地址
	ServerHost string `yaml:"ServerHost"`
	// LoginURL 登陆授权地址
	LoginURL string `yaml:"LoginURL"`
}

// Violet ...
type Violet struct {
	Config Config
}

// NewViolet ...
func NewViolet(c Config) Violet {
	return Violet{
		Config: c,
	}
}

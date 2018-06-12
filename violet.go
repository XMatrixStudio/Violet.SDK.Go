package violetSdk

type Config struct {
	ClientID string `yaml:"ClientID"`
	// ClientKey KEY
	ClientKey string `yaml:"ClientKey"`
	// ServerHost 服务器地址
	ServerHost string `yaml:"ServerHost"`
	// LoginURL 登陆授权地址
	LoginURL string `yaml:"LoginURL"`
}

type Violet struct {
	Config Config
}

func NewViolet(c Config) Violet {
	return Violet{
		Config: c,
	}
}

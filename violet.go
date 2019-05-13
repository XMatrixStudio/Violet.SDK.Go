package violet

// Config 配置文件
type Config struct {
	// ClientID 应用ID
	ClientID string
	// ClientKey 应用地址
	ClientKey string
	// ServerHost 服务器地址
	ServerHost string
}

// Violet ...
type Violet struct {
	Config Config
}

// NewViolet ...
func NewViolet(c Config) *Violet {
	return &Violet{
		Config: c,
	}
}

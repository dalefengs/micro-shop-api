package config

// Config 配置信息
type Config struct {
	Name     string       `mapstructure:"name" json:"name"`
	Host     string       `mapstructure:"host" json:"host"`
	Port     int          `mapstructure:"port" json:"port"`
	Services services     `mapstructure:"services" json:"services"`
	JwtInfo  JWTConfig    `mapstructure:"jwt" json:"jwt"`
	Redis    redisConfig  `mapstructure:"redis" json:"redis"`
	Consul   consulConfig `mapstructure:"consul" json:"consul"`
	Alipay   alipayConfig `mapstructure:"alipay" json:"alipay"` // 支付宝支付配置

}

type services struct {
	UserOpSrvName string `mapstructure:"userop_srv_name" json:"userop_srv_name"`
	GoodsSrvName  string `mapstructure:"goods_srv_name" json:"goods_srv_name"`
}

type alipayConfig struct {
	Appid string `mapstructure:"appid" json:"appid"`
	// 应用私钥
	AppPrivateKey string `mapstructure:"" json:"app_private_key"`
	// 支付回调URL
	NotifyUrl string `mapstructure:"notify_url" json:"notify_url"`
	// 支付成功跳转URL
	ReturnUrl string `mapstructure:"return_url" json:"return_url"`
	// 超时时间
	TimeoutExpress string `mapstructure:"timeout_express" json:"timeout_express"`

	// 证书在初始化时加载
	// 支付宝公钥证书
	AlipayPublicContentRSA2 []byte
	// 支付宝根证书
	AlipayRootContent []byte
	// 应用公钥证书
	AppPublicContent []byte
}

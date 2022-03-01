package config

// Config 配置信息
type Config struct {
	Name    string       `mapstructure:"name" json:"name"`
	Host    string       `mapstructure:"host" json:"host"`
	Port    int          `mapstructure:"port" json:"port"`
	JwtInfo JWTConfig    `mapstructure:"jwt" json:"jwt"`
	Consul  consulConfig `mapstructure:"consul" json:"consul"`
	Aliyun  aliyunConfig `mapstructure:"aliyun" json:"aliyun"`
}

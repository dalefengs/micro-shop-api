package config

// Config 配置信息
type Config struct {
	Name     string       `mapstructure:"name" json:"name"`
	Port     int          `mapstructure:"port" json:"port"`
	Services services     `mapstructure:"services" json:"services"`
	JwtInfo  JWTConfig    `mapstructure:"jwt" json:"jwt"`
	Aliyun   aliyun       `mapstructure:"aliyun" json:"aliyun"`
	Redis    redisConfig  `mapstructure:"redis" json:"redis"`
	Consul   consulConfig `mapstructure:"consul" json:"consul"`
}

type services struct {
	UserSrvName string `mapstructure:"user_srv_name" json:"user_srv_name"`
}

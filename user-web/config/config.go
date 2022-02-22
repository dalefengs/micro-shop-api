package config

// Config 配置信息
type Config struct {
	Name     string       `mapstructure:"name"`
	Port     int          `mapstructure:"port"`
	Services services     `mapstructure:"services"`
	JwtInfo  JWTConfig    `mapstructure:"jwt"`
	Aliyun   aliyun       `mapstructure:"aliyun"`
	Redis    redisConfig  `mapstructure:"redis"`
	Consul   consulConfig `mapstructure:"consul"`
}

type services struct {
	UserSrvName string `mapstructure:"user_srv_name"`
}

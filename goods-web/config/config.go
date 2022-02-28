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
}

type services struct {
	GoodsSrvName string `mapstructure:"goods_srv_name" json:"goods_srv_name"`
}

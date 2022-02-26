package config

type redisConfig struct {
	Host     string       `mapstructure:"host" json:"host"`
	Port     int          `mapstructure:"port" json:"port"`
	Password string       `mapstructure:"password" json:"password"`
	Prefix   PrefixConfig `mapstructure:"prefix" json:"prefix"`
}

// PrefixConfig Redis 前缀
type PrefixConfig struct {
	Users    string `mapstructure:"users" json:"users"`
	Register string `mapstructure:"register" json:"register"`
	Forget   string `mapstructure:"forget" json:"forget"`
}

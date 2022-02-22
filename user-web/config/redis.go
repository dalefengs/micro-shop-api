package config

type redisConfig struct {
	Host     string       `mapstructure:"host"`
	Port     int          `mapstructure:"port"`
	Password string       `mapstructure:"password"`
	Prefix   PrefixConfig `mapstructure:"prefix"`
}

// PrefixConfig Redis 前缀
type PrefixConfig struct {
	Users    string `mapstructure:"users"`
	Register string `mapstructure:"register"`
	Forget   string `mapstructure:"forget"`
}

package config

type redisConfig struct {
	Host     string       `mapstructure:"host" json:"host"`
	Port     int          `mapstructure:"port" json:"port"`
	Password string       `mapstructure:"password" json:"password"`
	Prefix   PrefixConfig `mapstructure:"prefix" json:"prefix"`
}

// PrefixConfig Redis 前缀
type PrefixConfig struct {
}

package config

type Nacos struct {
	Host          string `mapstructure:"host"`
	Port          uint64 `mapstructure:"port"`
	User          string `mapstructure:"user"`
	Password      string `mapstructure:"password"`
	Namespace     string `mapstructure:"name_space"`
	Group         string `mapstructure:"group"`
	UserApiDataId string `mapstructure:"user_api_data_id"`
}

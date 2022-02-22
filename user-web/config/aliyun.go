package config

// 阿里云
type aliyun struct {
	AccessKeyId     string `mapstructure:"accessKeyId"`
	AccessKeySecret string `mapstructure:"accessKeySecret"`
	Sms             sms    `mapstructure:"sms"`
}

// 阿里云短信
type sms struct {
	SignName     string `mapstructure:"signName"`
	TemplateCode string `mapstructure:"templateCode"`
	Expire       int    `mapstructure:"expire"`
}

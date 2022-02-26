package config

// 阿里云
type aliyun struct {
	AccessKeyId     string `mapstructure:"accessKeyId" json:"accessKeyId"`
	AccessKeySecret string `mapstructure:"accessKeySecret" json:"accessKeySecret"`
	Sms             sms    `mapstructure:"sms" json:"sms"`
}

// 阿里云短信
type sms struct {
	SignName     string `mapstructure:"signName" json:"signName"`
	TemplateCode string `mapstructure:"templateCode" json:"templateCode"`
	Expire       int    `mapstructure:"expire" json:"expire"`
}

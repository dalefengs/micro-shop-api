package config

// 阿里云
type aliyunConfig struct {
	AccessKeyId     string `mapstructure:"accessKeyId" json:"accessKeyId"`
	AccessKeySecret string `mapstructure:"accessKeySecret" json:"accessKeySecret"`
	Bucket          string `mapstructure:"bucket" json:"bucket"`
	// 一定要公网ip才能回调
	CallbackUrl string `mapstructure:"callbackUrl" json:"callbackUrl"`
	UploadDir   string `mapstructure:"uploadDir" json:"uploadDir"`
	ExpireTime  int64  `mapstructure:"expireTime" json:"expireTime"`
}

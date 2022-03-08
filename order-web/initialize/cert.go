package initialize

import (
	"go.uber.org/zap"
	"io/ioutil"
	"micro-shop-api/order-web/global"
	"micro-shop-api/order-web/utils"
	"os"
)

// InitCert 加载证书文件
// TODO 监听配置文件变化并重新加载
func InitCert() {
	path := utils.GetProjectPath(global.Config.Name) + "/config/cert/"
	RSA2, err := os.Open(path + "alipayCertPublicKey_RSA2.crt")
	if err != nil {
		zap.S().Fatalf("打开 alipayCertPublicKey_RSA2.crt 失败:%s", err.Error())
	}
	rootKey, err := os.Open(path + "alipayRootCert.crt")
	if err != nil {
		zap.S().Fatalf("打开 alipayRootCert.crt 失败:%s", err.Error())

	}
	appPublicKey, err := os.Open(path + "appCertPublicKey.crt")
	if err != nil {
		zap.S().Fatalf("打开 appCertPublicKey.crt 失败:%s", err.Error())
	}
	global.Config.Alipay.AlipayPublicContentRSA2, err = ioutil.ReadAll(RSA2)
	if err != nil {
		zap.S().Fatalf("读取 appPublicKey.crt 失败:%s", err.Error())
	}
	global.Config.Alipay.AlipayRootContent, err = ioutil.ReadAll(rootKey)
	if err != nil {
		zap.S().Fatalf("读取 appPublicKey.crt 失败:%s", err.Error())
	}
	global.Config.Alipay.AppPublicContent, err = ioutil.ReadAll(appPublicKey)
	if err != nil {
		zap.S().Fatalf("读取 appPublicKey.crt 失败:%s", err.Error())
	}

}

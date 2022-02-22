package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"math/rand"
	"strings"
	"time"
)

// IsDev 是否是本地环境
// 开发环境需要配置环境变量 export MICRO_SHOP_DEV=true
// Windows 重启编辑器  Linux 注销登录
func IsDev() bool {
	viper.AutomaticEnv()
	return viper.GetBool("MICRO_SHOP_DEV")
}

// RemoveToStruct 删除错误信息键值的struct信息
func RemoveToStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// GetFirstError 获取 map 中一个元素(不保证顺序性)
func GetFirstError(fields map[string]string) string {
	for _, v := range fields {
		return v
	}
	return ""
}

// GenValidateCode 生成随机数验证码
func GenValidateCode(codeLen int) string {
	numeric := [10]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	nr := len(numeric)
	rand.Seed(time.Now().UnixMicro()) // 随机数种子
	var code strings.Builder
	for i := 0; i < codeLen; i++ {
		_, _ = fmt.Fprintf(&code, "%d", numeric[rand.Intn(nr)])
	}
	return code.String()
}

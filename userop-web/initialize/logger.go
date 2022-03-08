package initialize

import "go.uber.org/zap"

func InitZapLogger() {
	logger, err := zap.NewDevelopment()
	//logger, err := zap.NewProduction() // 生产环境
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}

package initialize

import (
	"github.com/edynt/go-ecommerce-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()

	global.Logger.Info("Config log ok!!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()
	InitRouter()

	r := InitRouter()
	r.Run()
}

package initialize

import (
	"github.com/edynt/go-ecommerce-api/global"
	"github.com/edynt/go-ecommerce-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}

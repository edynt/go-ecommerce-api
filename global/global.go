package global

import (
	"github.com/edynt/go-ecommerce-api/pkg/logger"
	"github.com/edynt/go-ecommerce-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)

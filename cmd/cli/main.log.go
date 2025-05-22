package main

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age: %d", "Edynt", 20)

	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipsGo"), zap.Int("age", 40))

	// logger := zap.NewExample()
	// logger.Info("New Example")

	// loggerDev, _ := zap.NewDevelopment()
	// loggerDev.Info("New Development")

	// loggerProd, _ := zap.NewProduction()
	// loggerProd.Info("New Production")

	err := os.MkdirAll("./log", os.ModePerm)
	if err != nil {
		log.Fatalf("Không thể tạo thư mục: %v", err)
	}

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zap.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))

	logger.Error("Error log", zap.Int("line", 2))
}

// format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// timestamp => dd/mm/yyyy
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> time
	encodeConfig.TimeKey = "time"

	// level
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// caller
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

// write
func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

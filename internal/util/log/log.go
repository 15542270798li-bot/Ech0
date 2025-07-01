package util

import (
	model "github.com/lin-snow/ech0/internal/model/common"
	"go.uber.org/zap"
)

// Logger 全局日志记录器
var Logger *zap.Logger

// InitLogger 初始化日志记录器
func InitLogger() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(model.INIT_LOGGER_PANIC + ": " + err.Error())
	}
}

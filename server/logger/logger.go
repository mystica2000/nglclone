package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger;

func InitLogger() {
	Log, _ = zap.NewDevelopment()
	defer Log.Sync()
}
package core

import (
	"log"
	"os"
	"server/global"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.Logger {
	zapconf := global.Config.Zap
	writeSync := getLogWriter(zapconf.Filename, zapconf.MaxSize, zapconf.MaxBackups, zapconf.MaxAge, zapconf.IsConsolePrint)
	if zapconf.IsConsolePrint {
		writeSync = zapcore.NewMultiWriteSyncer(writeSync, zapcore.AddSync(os.Stdout))
	}
	var logLevel zapcore.Level
	err := logLevel.UnmarshalText([]byte(zapconf.Level))
	if err != nil {
		log.Fatalf("日志级别设置错误: %v", err)
	}
	encoder := getIncoder()
	core := zapcore.NewCore(encoder, writeSync, logLevel)
	logger := zap.New(core, zap.AddCaller())

	return logger
}

func getLogWriter(filename string, maxsize int, maxbackups int, maxage int, isconsoleprint bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxsize,
		MaxBackups: maxbackups,
		MaxAge:     maxage,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getIncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

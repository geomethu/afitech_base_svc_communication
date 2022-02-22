package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

//type Logger struct{ *zap.Logger }

func InitLogger()  {
	writeSync := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSync, zapcore.DebugLevel)

	Logger = zap.New(core)
	//return &Logger{zap.New(core)}
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(config)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("/home/methu/Hub/golang/2_22/AFITECH_BASE/test.log")
	return zapcore.AddSync(file)
}

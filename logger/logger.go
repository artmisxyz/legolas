package logger

import (
	"fmt"
	"github.com/artmisxyz/legolas/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func New(conf configs.Config) *zap.Logger {
	var lvl zapcore.Level
	err := lvl.Set(conf.Logger.Level)
	if err != nil {
		fmt.Printf("cannot parse log level %s: %s", conf.Logger.Level, err)
		lvl = zapcore.WarnLevel
	}
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	defaultCore := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stderr)), lvl)
	cores := []zapcore.Core{
		defaultCore,
	}

	core := zapcore.NewTee(cores...)
	return zap.New(core, zap.AddCaller())
}

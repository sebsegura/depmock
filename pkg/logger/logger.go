package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func NewLogger() *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "ts",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(time.RFC3339))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	return zap.New(core)
}

type key int

const (
	_key key = iota
)

func NewContextWithLogger(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, _key, log)
}

func Logger(ctx context.Context) *zap.Logger {
	log, _ := ctx.Value(_key).(*zap.Logger)
	return log
}

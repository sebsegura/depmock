package lambda

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type Middleware[I, O any] func(next HandlerFunc[I, O]) HandlerFunc[I, O]

func LoggerMiddleware[I, O any](next HandlerFunc[I, O]) HandlerFunc[I, O] {
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
	log := zap.New(core).Sugar()

	return func(ctx context.Context, in *I) (*O, error) {
		ctx = context.WithValue(ctx, "logger", log)
		log.With(zap.Any("event", in)).Info("starting execution...")

		out, err := next(ctx, in)
		if err != nil {
			log.With(zap.Any("error", err)).Error("execution has an error")
		} else {
			log.With(zap.Any("response", out)).Info("success")
		}

		return out, err
	}
}

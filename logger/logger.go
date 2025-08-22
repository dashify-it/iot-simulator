package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func InitLogger(debug bool) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,        // INFO in blue, WARN in yellow, etc.
		EncodeTime:     zapcore.TimeEncoderOfLayout("15:04:05"), // HH:MM:SS
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	var Level zapcore.Level
	if debug {
		Level = zap.DebugLevel
	} else {
		Level = zap.InfoLevel
	}
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), Level)
	Log = zap.New(core, zap.AddCaller()).Sugar()

}

// Sync should be called before the app exits
func Sync() {
	_ = Log.Sync()
}

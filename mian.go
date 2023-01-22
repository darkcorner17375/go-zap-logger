package main

import (
	"go-zap-logger/log/logger"
)

func main() {
	log1 := logger.New()
	log1.Config.SetProjectName("log1")
	log1.Error("msg log1")
	log1.Config.SetJSONFormat(false)
	log1.Error("msg log1")
	log1.ApplyConfig()
	log1.Error("msg log1")

	// log2 := logger.New()
	// log2.Config.SetProjectName("log2")
	// log2.ApplyConfig()
	// log2.Error("msg log2")
}

// package main

// import (
//  "os"

//  "go.uber.org/zap"
//  "go.uber.org/zap/zapcore"
// )

// func NewCustomEncoderConfig() zapcore.EncoderConfig {
//  return zapcore.EncoderConfig{
//   TimeKey:        "ts",
//   LevelKey:       "level",
//   NameKey:        "logger",
//   CallerKey:      "caller",
//   FunctionKey:    zapcore.OmitKey,
//   MessageKey:     "msg",
//   StacktraceKey:  "stacktrace",
//   LineEnding:     zapcore.DefaultLineEnding,
//   EncodeLevel:    zapcore.CapitalColorLevelEncoder,
//   EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
//   EncodeDuration: zapcore.SecondsDurationEncoder,
//   EncodeCaller:   zapcore.ShortCallerEncoder,
//  }
// }

// func main() {
//  atom := zap.NewAtomicLevelAt(zap.DebugLevel)
//  core := zapcore.NewCore(
//   zapcore.NewConsoleEncoder(NewCustomEncoderConfig()),
//   zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
//   atom,
//  )
//  logger := zap.New(core, zap.AddCaller(), zap.Development())
//  defer logger.Sync()

//  // 配置 zap 包的全局變量
//  zap.ReplaceGlobals(logger)

//  // 運行時安全地更改 logger 日記級別
//  atom.SetLevel(zap.InfoLevel)
//  sugar := logger.Sugar()
//  // 問題 1: debug 級別的日誌打印到控制檯了嗎?
//  sugar.Debug("debug")
//  sugar.Info("info")
//  sugar.Warn("warn")
// //  sugar.DPanic("dPanic")
//  // 問題 2: 最後的 error 會打印到控制檯嗎?
//  sugar.Error("error")
// }

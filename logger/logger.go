package logger

//go get -u go.uber.org/zap
//go get -u github.com/natefinch/lumberjack

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//預設Logger
type Logger struct {
	Config *Config
	logger *zap.SugaredLogger
}

//自定義Logger
func New() *Logger {
	logger := &Logger{
		Config: newConfig(),
	}
	logger.SubmitConfig()
	return logger
}

//確定Config設定
func (l *Logger) SubmitConfig() {
	logConf := l.Config
	cores := []zapcore.Core{}
	nowTime := time.Now()
	nowTimeStr := nowTime.Format("2006-01-02")

	var encoder zapcore.Encoder

	if logConf.JsonFormat {
		encoder = zapcore.NewJSONEncoder(getConsoleEncoder())
	} else {
		encoder = zapcore.NewConsoleEncoder(getConsoleEncoder())
	}

	logConf.AtomicLevel.SetLevel(getLevel(logConf.DefaultLevel))

	if logConf.ConsoleOut {
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(encoder, writer, logConf.AtomicLevel)
		cores = append(cores, core)
	}

	if logConf.FileOut {
		encoder = zapcore.NewJSONEncoder(getFileEncoder())
		writeSyncer := getLogWriter(nowTimeStr)
		core := zapcore.NewCore(encoder, writeSyncer, logConf.AtomicLevel)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	logger := zap.New(combinedCore,
		zap.AddCallerSkip(logConf.CallerSkip),
		zap.AddStacktrace(getLevel(logConf.StacktraceLevel)),
		zap.AddCaller(),
	)

	if logConf.ProjectName != "" {
		logger = logger.Named(logConf.ProjectName)
	}

	defer logger.Sync()

	l.logger = logger.Sugar()
}

//Debug Debug log
func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

//Debugf Debug format log
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

//Debugw Debugw log
func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

//Info Info log
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

//Infof Info format log
func (l *Logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

//Infow Infow log
func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

//Warn Warn log
func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

//Warnf Warn format log
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

//Warnw Warnw log
func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

//Error Error log
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

//Errorf Error format log
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

//Errorw Errorw log
func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

//Panic Panic log
func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

//Panicf Panic format log
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

//Panicw Panicw log
func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues...)
}

//Fatal Fatal log
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

//Fatalf Fatal format log
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

//Fatalw Fatalw log
func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func getConsoleEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		LevelKey:       "Level",
		TimeKey:        "Time",
		MessageKey:     "Message",
		NameKey:        "Project",
		CallerKey:      "Caller",
		StacktraceKey:  "Trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func getFileEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		LevelKey:       "Level",
		TimeKey:        "Time",
		MessageKey:     "Message",
		NameKey:        "Project",
		CallerKey:      "Caller",
		StacktraceKey:  "Trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func getLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func getLogWriter(fileName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/" + fileName + ".log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func (l *Logger) GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		l.Config.SetProjectName("gin")
		l.SubmitConfig()

		cost := time.Since(start)
		l.logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

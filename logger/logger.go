package logger

//go get -u go.uber.org/zap
//go get -u github.com/natefinch/lumberjack

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*base Logger 基本Logger*/
type Logger struct {
	Config *Config
	logger *zap.SugaredLogger
}

/*Customer Logger 自定義Logger*/
func New() *Logger {
	logger := &Logger{
		Config: newConfig(),
	}
	logger.ApplyConfig()
	return logger
}

/*ApplyConfig 应用当前Config配置*/
func (l *Logger) ApplyConfig() {
	conf := l.Config
	cores := []zapcore.Core{}

	var encoder zapcore.Encoder

	if conf.jsonFormat {
		encoder = zapcore.NewJSONEncoder(getEncoder())
	} else {
		encoder = zapcore.NewConsoleEncoder(getEncoder())
	}

	conf.atomicLevel.SetLevel(getLevel(conf.defaultLogLevel))

	if conf.consoleOut {
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(encoder, writer, conf.atomicLevel)
		cores = append(cores, core)
	}

	if conf.fileOut {
		writeSyncer := getLogWriter()
		core := zapcore.NewCore(encoder, writeSyncer, conf.atomicLevel)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	logger := zap.New(combinedCore,
		zap.AddCallerSkip(conf.callerSkip),
		zap.AddStacktrace(getLevel(conf.stacktraceLevel)),
		zap.AddCaller(),
	)

	if conf.projectName != "" {
		logger = logger.Named(conf.projectName)
	}

	defer logger.Sync()

	l.logger = logger.Sugar()
}

/*Debug Debug log*/
func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

/*Debugf Debug format log*/
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

/*Debugw Debugw log*/
func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

/*Info Info log*/
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

/*Infof Info format log*/
func (l *Logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

/*Infow Infow log*/
func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

/*Warn Warn log*/
func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

/*Warnf Warn format log*/
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

/*Warnw Warnw log*/
func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

/*Error Error log*/
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

/*Errorf Error format log*/
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

/*Errorw Errorw log*/
func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

/*Panic Panic log*/
func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

/*Panicf Panic format log*/
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

/*Panicw Panicw log*/
func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues...)
}

/*Fatal Fatal log*/
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

/*Fatalf Fatal format log*/
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

/*Fatalw Fatalw log*/
func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

// func getFileWriter(path, name string, rotationTime, rotationCount uint) io.Writer {
// 	writer, err := rotatelogs.New(
// 		filepath.Join(path, name+".%Y%m%d%H.log"),
// 		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour), // 日志切割时间间隔
// 		rotatelogs.WithRotationCount(rotationCount),                        // 文件最大保存份数
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return writer
// }

func getEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		LevelKey:       "L",
		TimeKey:        "T",
		MessageKey:     "M",
		NameKey:        "N",
		CallerKey:      "C",
		StacktraceKey:  "S",
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

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

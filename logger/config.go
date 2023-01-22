package logger

import (
	"go.uber.org/zap"
)

//Logger設置
type Config struct {
	DefaultLevel    string          //日誌紀錄等級
	StacktraceLevel string          //最高紀錄等級
	AtomicLevel     zap.AtomicLevel //動態更動等級
	ProjectName     string          //專案名稱
	CallerSkip      int             //CallerSkip
	JsonFormat      bool            //输出json格式
	ConsoleOut      bool            //是否输出到console
	FileOut         bool
}

//建立文檔設置
type LogWriterConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

//預設Logger設置
func newConfig() *Config {
	return &Config{
		DefaultLevel:    "info",
		StacktraceLevel: "panic",
		AtomicLevel:     zap.NewAtomicLevel(),
		ProjectName:     "",
		CallerSkip:      1,
		JsonFormat:      false,
		ConsoleOut:      true,
		FileOut:         true,
	}
}

//設置日誌紀錄等級
func (c *Config) SetLevel(level string) {
	c.AtomicLevel.SetLevel(getLevel(level))
}

//設置堆棧日誌等級
func (c *Config) SetStacktraceLevel(level string) *Config {
	c.StacktraceLevel = level
	return c
}

//設置專案名稱
func (c *Config) SetProjectName(projectName string) *Config {
	c.ProjectName = projectName
	return c
}

//設置callerSkip次數
func (c *Config) SetCallerSkip(callerSkip int) *Config {
	c.CallerSkip = callerSkip
	return c
}

//設置是否JSON輸出
func (c *Config) SetJSONFormat(enable bool) *Config {
	c.JsonFormat = enable
	return c
}

//設置是否Console輸出
func (c *Config) SetConsoleOut(enable bool) *Config {
	c.ConsoleOut = enable
	return c
}

//設置寫入文件設定
func (l *LogWriterConfig) SetLogWriter(path string, maxSize, maxBackups, maxAge int, compress bool) *LogWriterConfig {
	l.Filename = path
	l.MaxSize = maxSize
	l.MaxBackups = maxBackups
	l.MaxAge = maxAge
	l.Compress = compress
	return l
}

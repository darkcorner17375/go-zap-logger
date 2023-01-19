package logger

import (
	"go.uber.org/zap"
)

//新增Logger設定 結構體
type Config struct {
	defaultLogLevel string          //日誌紀錄等級
	stacktraceLevel string          //最高紀錄等級
	atomicLevel     zap.AtomicLevel //動態更動等級
	projectName     string          //專案名稱
	callerSkip      int             //CallerSkip
	jsonFormat      bool            //输出json格式
	consoleOut      bool            //是否输出到console
	fileOut         bool
}

type LogWriter struct {
}

//預設自定義Logger配置
func newConfig() *Config {
	return &Config{
		defaultLogLevel: "info",
		stacktraceLevel: "panic",
		atomicLevel:     zap.NewAtomicLevel(),
		projectName:     "",
		callerSkip:      1,
		jsonFormat:      true,
		consoleOut:      true,
		fileOut:         true,
	}
}

/*SetLevel 设置日志记录级别*/
func (c *Config) SetLevel(level string) {
	c.atomicLevel.SetLevel(getLevel(level))
}

/*SetStacktraceLevel 设置堆栈跟踪的日志级别*/
func (c *Config) SetStacktraceLevel(level string) {
	c.stacktraceLevel = level
}

/*SetProjectName 设置ProjectName*/
func (c *Config) SetProjectName(projectName string) {
	c.projectName = projectName
}

/*SetCallerSkip 设置callerSkip次数*/
func (c *Config) SetCallerSkip(callerSkip int) {
	c.callerSkip = callerSkip
}

/*EnableJSONFormat 开启JSON格式化输出*/
func (c *Config) EnableJSONFormat() {
	c.jsonFormat = true
}

/*DisableJSONFormat 关闭JSON格式化输出*/
func (c *Config) DisableJSONFormat() {
	c.jsonFormat = false
}

/*EnableConsoleOut 开启Console输出*/
func (c *Config) EnableConsoleOut() {
	c.consoleOut = true
}

/*DisableConsoleOut 关闭Console输出*/
func (c *Config) DisableConsoleOut() {
	c.consoleOut = false
}

// File:    logging
// Version: 1.0.0
// Creator: JoeLang
// Date:    2020/8/31 23:04
// DESC:

package logger

import (
	"sync"
	"time"
)

const (
	LevelEmergency     = iota // 系统级紧急，比如磁盘出错，内存异常，网络不可用等
	LevelAlert                // 系统级警告，比如数据库访问异常，配置文件出错等
	LevelCritical             // 系统级危险，比如权限出错，访问异常等
	LevelError                // 用户级错误
	LevelWarning              // 用户级警告
	LevelInformational        // 用户级信息
	LevelDebug                // 用户级调试
	LevelTrace                // 用户级基本输出
)

var LevelMap = map[string]int{
	"EMER": LevelEmergency,
	"ALRT": LevelAlert,
	"CRIT": LevelCritical,
	"ERRO": LevelError,
	"WARN": LevelWarning,
	"INFO": LevelInformational,
	"DEBG": LevelDebug,
	"TRAC": LevelTrace,
}

type logInfo struct {
	Time    string
	Level   string
	Path    string
	Name    string
	Content string
}

type nameLogger struct {
	Logger
	name   string
	config string
}

type LocalLogger struct {
	lock       sync.Mutex
	init       bool
	outputs    []*nameLogger
	appName    string
	callDepth  int
	timeFormat string
	usePath    string
}

// 默认日志输出
var defaultLogger *LocalLogger

// 注册实现的适配器， 当前支持控制台，文件和网络输出
var adapters = make(map[string]Logger)


// log provider interface
type Logger interface {
	Init(config string) error
	LogWrite(when time.Time, msg interface{}, level int) error
	Destroy()
}

// 日志输出适配器注册，log需要实现Init，LogWrite，Destroy方法
func Register(name string, log Logger) {
	if log == nil {
		panic("logs: Register provide is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("logs: Register called twice for provider " + name)
	}
	adapters[name] = log
}
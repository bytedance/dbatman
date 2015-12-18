// Copyright (c) All Rights Reserved
// @file    logger.go
// @author  王靖 (wangjild@gmail.com)
// @date    14-11-25 20:02:50
// @version $Revision: 1.0 $
// @brief

package log

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// SysLog 系统Log
var SysLog *ProxyLogger = nil

// AppLog 应用Log
var AppLog *ProxyLogger = nil

// Logger the log.Logger wrapper
type ProxyLogger struct {
	l *Logger
}

func logidGenerator() string {
	if i, err := rand.Int(rand.Reader, big.NewInt(1<<30-1)); err != nil {
		return "0"
	} else {
		return i.String()
	}
}

func comMessage(strfmt string, args ...interface{}) map[string]string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "?"
		line = 0
	}
	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}
	ret := map[string]string{
		"file": filepath.Base(file) + ":" + strconv.Itoa(line),
		"func": fnName,
		"msg":  fmt.Sprintf(strfmt, args...),
	}

	return ret
}

// Notice print notice message to logfile
func (lg *ProxyLogger) Notice(strfmt string, args ...interface{}) {
	lg.l.Notice(comMessage(strfmt, args...), logidGenerator())
}

// Debug print debug message to logfile
func (lg *ProxyLogger) Debug(strfmt string, args ...interface{}) {
	lg.l.Debug(comMessage(strfmt, args...), logidGenerator())
}

// Warn print warning message to logfile
func (lg *ProxyLogger) Warn(strfmt string, args ...interface{}) {
	lg.l.Warn(comMessage(strfmt, args...), logidGenerator())
}

// Fatal print fatal message to logfile
func (lg *ProxyLogger) Fatal(strfmt string, args ...interface{}) {
	lg.l.Fatal(comMessage(strfmt, args...), logidGenerator())
}

// Config Config of One Log Instance
type Config struct {
	FilePath string
	LogLevel int
	AppTag   string
}

func init() {
	realInit(&Config{FilePath: "/dev/stdout", LogLevel: 0},
		&Config{FilePath: "/dev/stdout", LogLevel: 0})
}

var once sync.Once

func Init(syslog, applog *Config) {
	f := func() {
		realInit(syslog, applog)
	}
	once.Do(f)
}

func realInit(syslog, applog *Config) {
	SysLog = &ProxyLogger{
		l: NewLogger(syslog.FilePath),
	}
	SysLog.l.SetLevel(syslog.LogLevel)
	SysLog.l.SetAppTag(defaultAppTag())

	AppLog = &ProxyLogger{
		l: NewLogger(applog.FilePath),
	}
	AppLog.l.SetLevel(applog.LogLevel)
	AppLog.l.SetAppTag(defaultAppTag())
}

func defaultAppTag() string {
	return "mysql-proxy"
}

/* vim: set expandtab ts=4 sw=4 */

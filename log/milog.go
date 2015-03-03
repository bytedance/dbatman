package log

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	LevelTrace = iota
	LevelDebug
	LevelNotice
	LevelWarn
	LevelFatal
)

const (
	DefaultKey    = "DefaultKey"
	DefaultLogId  = "000000000000"
	DefaultAppTag = "DefaultAppTag"
)

// map TRACE, DEBUG, NOTICE, WARN, FATAL to 0, 1, 2, 3, 4
var (
	Level = []string{"TRACE", "DEBUG", "NOTICE", "WARN", "FATAL"}
)

type Logger struct {
	logfd    *os.File
	level    int
	apptag   string
	hostname string
	lock     *sync.Mutex
}

// NewLogger return a Logger instance,
// the params is filename and apptag is optional
func NewLogger(filename string, apptag ...string) *Logger {
	var realAppTag string
	if len(apptag) == 0 {
		realAppTag = DefaultAppTag
	} else {
		realAppTag = apptag[0]
	}

	// panic when can not get hostname
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	// panic when can not open log file
	logfd, err := os.OpenFile(
		filename,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}

	return &Logger{
		logfd:    logfd,
		level:    LevelNotice,
		apptag:   realAppTag,
		hostname: hostname,
		lock:     new(sync.Mutex),
	}
}

// SetLevel set the log level, default is LevelNotice
func (this *Logger) SetLevel(level int) {
	this.level = level
}

// SetAppTag  set log's apptag
func (this *Logger) SetAppTag(apptag string) {
	this.apptag = apptag
}

// Level return this logger's level
func (this *Logger) Level() int {
	return this.level
}

// AppTag return this logger's apptag
func (this *Logger) AppTag() string {
	return this.apptag
}

// It should be locked while calling write method.
func (this *Logger) write(msg string) error {
	if !strings.HasSuffix(msg, "\n") {
		msg = msg + "\n"
	}

	this.lock.Lock()
	defer this.lock.Unlock()
	_, err := this.logfd.WriteString(msg)
	return err
}

// Flush will write all logs from os's buffer to disk
func (this *Logger) Flush() {
	this.logfd.Sync()
	this.logfd.Close()
}

func (this *Logger) Trace(v interface{}, logid ...string) error {
	if !this.suitLevel(LevelTrace) {
		return nil
	}
	return this.write(this.format(LevelTrace, v, logid))
}

func (this *Logger) Debug(v interface{}, logid ...string) error {
	if !this.suitLevel(LevelDebug) {
		return nil
	}
	return this.write(this.format(LevelDebug, v, logid))
}

func (this *Logger) Notice(v interface{}, logid ...string) error {
	if !this.suitLevel(LevelNotice) {
		return nil
	}
	return this.write(this.format(LevelNotice, v, logid))
}

func (this *Logger) Warn(v interface{}, logid ...string) error {
	if !this.suitLevel(LevelWarn) {
		return nil
	}
	return this.write(this.format(LevelWarn, v, logid))
}

func (this *Logger) Fatal(v interface{}, logid ...string) error {
	if !this.suitLevel(LevelFatal) {
		return nil
	}
	return this.write(this.format(LevelFatal, v, logid))
}

func (this *Logger) suitLevel(level int) bool {
	if level < this.level {
		return false
	}
	return true
}

// format generate a standard line of log
func (this *Logger) format(level int, v interface{}, logid []string) string {
	var id string
	if len(logid) > 0 {
		id = logid[0]
	} else {
		id = DefaultLogId
	}

	prefix := ""
	var logTuples = []string{
		time.Now().Format("2006-01-02 15:04:05"),
		this.apptag,
		this.hostname,
		Level[level],
		id,
	}

	for _, item := range logTuples {
		prefix += "[" + item + "] "
	}

	var (
		body []byte
		err  error
	)
	body, err = json.Marshal(v)
	if err != nil {
		body, _ = json.Marshal(
			map[string]interface{}{
				DefaultKey: v,
			},
		)
	}

	return prefix + string(body)
}

// StdContent used to store temporary log content
type StdContent struct {
	data   map[string]interface{}
	logger *Logger
	lock   *sync.Mutex
}

// NewStdContent return an temporary StdContent
func (this *Logger) NewStdContent() *StdContent {
	return &StdContent{
		data:   make(map[string]interface{}),
		logger: this,
		lock:   &sync.Mutex{},
	}
}

// SetVal add (key, value) pair to log
func (sc *StdContent) SetVal(key string, val interface{}) {
	sc.lock.Lock()
	sc.data[key] = val
	sc.lock.Unlock()
}

func (sc *StdContent) Trace(logid ...string) error {
	if !sc.logger.suitLevel(LevelTrace) {
		return nil
	}
	formatStr := sc.logger.format(LevelTrace, sc.data, logid)
	return sc.logger.write(formatStr)
}

func (sc *StdContent) Debug(logid ...string) error {
	if !sc.logger.suitLevel(LevelDebug) {
		return nil
	}
	formatStr := sc.logger.format(LevelDebug, sc.data, logid)
	return sc.logger.write(formatStr)
}

func (sc *StdContent) Notice(logid ...string) error {
	if !sc.logger.suitLevel(LevelNotice) {
		return nil
	}
	formatStr := sc.logger.format(LevelNotice, sc.data, logid)
	return sc.logger.write(formatStr)
}

func (sc *StdContent) Warn(logid ...string) error {
	if !sc.logger.suitLevel(LevelWarn) {
		return nil
	}
	formatStr := sc.logger.format(LevelWarn, sc.data, logid)
	return sc.logger.write(formatStr)
}

func (sc *StdContent) Fatal(logid ...string) error {
	if !sc.logger.suitLevel(LevelFatal) {
		return nil
	}
	formatStr := sc.logger.format(LevelFatal, sc.data, logid)
	return sc.logger.write(formatStr)
}

package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := NewLogger("test.log", "testtag")
	for i := 0; i <= LevelFatal; i++ {
		logger.SetLevel(i)
		outs(logger)
	}

	logger.Flush()
}

func outs(logger *Logger) {
	logger.Trace("trace")
	logger.Debug("debug")
	logger.Notice("notice")
	logger.Warn("warn")
	logger.Fatal("fatal")
}

func TestSetLevel(t *testing.T) {
	logger := NewLogger("test.log", "testtag")
	logger.SetAppTag("testtag")
	_ = logger.Level()
	_ = logger.AppTag()
}

func TestLevel(t *testing.T) {
	logger := NewLogger("test.log", "testtag")

	levels := []string{"Trace", "Debug", "Notice", "Warn", "Fatal"}

	levelMap := map[string]int{
		"Trace":  LevelTrace,
		"Debug":  LevelDebug,
		"Notice": LevelNotice,
		"Warn":   LevelWarn,
		"Fatal":  LevelFatal,
	}
	for i, level := range levels {
		logger.SetLevel(levelMap[level])
		if logger.level != i {
			t.Error("Test Log SetLevel Error")
		}
	}
	logger.Flush()
	t.Log("Test Level Pass")
}

func TestLogContentLog(t *testing.T) {
	logger := NewLogger("test.log")
	logger.SetLevel(LevelDebug)

	sc := logger.NewStdContent()
	t.Log("Test NewStdContent With Id Pass")

	mkContent(sc)

	scNoId := logger.NewStdContent()
	t.Log("Test NewStdContent Without Id Pass")
	mkContent(scNoId)

	for i := 0; i <= LevelFatal; i++ {
		logger.SetLevel(i)
		stdOuts(sc, "LogIddddddd")
		stdOuts(scNoId)
	}
	logger.Flush()
}

func stdOuts(sc *StdContent, logId ...string) {
	sc.Trace(logId...)
	sc.Debug(logId...)
	sc.Notice(logId...)
	sc.Warn(logId...)
	sc.Fatal(logId...)
}

func BenchmarkLog(b *testing.B) {
	logger := NewLogger("test.log")
	logger.SetLevel(LevelDebug)
	for i := 0; i < b.N; i++ {
		sc := logger.NewStdContent()
		mkBenchContent(sc)
		sc.Trace("LogIddddddd")
		sc.Debug("LogIddddddd")
		sc.Notice("LogIddddddd")
		sc.Warn("LogIddddddd")
		sc.Fatal("LogIddddddd")
	}
	logger.Flush()
}

func mkBenchContent(sc *StdContent) {
	sc.SetVal("key0", "val0")
	sc.SetVal("nullval", nil)
	sc.SetVal("123", 123)
}

func mkContent(sc *StdContent) {
	sc.SetVal("key0", "val0")
	sc.SetVal("arr", []string{"val0", "val1", "val2"})
	sc.SetVal("hash", map[string]string{
		"v0": "val0",
		"v1": "val1",
		"v2": "val2",
	})
	sc.SetVal("nullval", nil)
	var interfaceVal interface{}
	interfaceVal = []interface{}{"123", 123, map[string]string{"v4": "val4"}}
	sc.SetVal("interface", interfaceVal)
	sc.SetVal("123", 123)
}

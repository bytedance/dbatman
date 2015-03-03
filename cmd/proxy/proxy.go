package main

import (
	"flag"
	"github.com/wangjild/go-mysql-proxy/config"
	. "github.com/wangjild/go-mysql-proxy/log"
	"github.com/wangjild/go-mysql-proxy/proxy"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var configFile *string = flag.String("config", "etc/proxy.yaml", "go mysql proxy config file")
var logLevel *int = flag.Int("loglevel", 0, "0-debug| 1-notice|2-warn|3-fatal")
var logFile *string = flag.String("logfile", "log/proxy.log", "go mysql proxy logfile")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	if len(*configFile) == 0 {
		SysLog.Fatal("must use a config file")
		return
	}

	cfg, err := config.ParseConfigFile(*configFile)
	if err != nil {
		SysLog.Fatal(err.Error())
		return
	}

	//Init(&Config{FilePath: *logFile, LogLevel: *logLevel},
	//	&Config{FilePath: *logFile, LogLevel: *logLevel})

	var svr *proxy.Server
	svr, err = proxy.NewServer(cfg)
	if err != nil {
		SysLog.Fatal(err.Error())
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		http.ListenAndServe(":11888", nil)
	}()

	go func() {
		sig := <-sc
		SysLog.Notice("Got signal [%d] to exit.", sig)
		svr.Close()
	}()

	svr.Run()
}

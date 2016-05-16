package main

import (
	"flag"
	"github.com/bytedance/dbatman/config"
	"github.com/bytedance/dbatman/database/cluster"
	"github.com/bytedance/dbatman/proxy"
	"github.com/ngaut/log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	configFile *string = flag.String("config", "etc/proxy.yaml", "go mysql proxy config file")
	logLevel   *int    = flag.Int("loglevel", 0, "0-debug| 1-notice|2-warn|3-fatal")
	logFile    *string = flag.String("logfile", "log/proxy.log", "go mysql proxy logfile")
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	if len(*configFile) == 0 {
		log.Fatal("must use a config file")
		os.Exit(1)
	}

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	if err = cluster.Init(cfg); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	go func() {
		cfg.CheckConfigUpdate(config.NotifyChan)
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	var svr *proxy.Server
	svr, err = proxy.NewServer(cfg)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	go func() {
		http.ListenAndServe(":11888", nil)
	}()

	go func() {
		sig := <-sc
		log.Infof("Got signal [%d] to exit.", sig)
		svr.Close()
	}()

	svr.Serve()
}

package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/bytedance/dbatman/config"
	"github.com/bytedance/dbatman/database/cluster"
	"github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/proxy"
	"github.com/ngaut/log"
)

var (
	configFile *string = flag.String("config", "etc/proxy.yaml", "go mysql proxy config file")
	logLevel   *int    = flag.Int("loglevel", 0, "0-debug| 1-notice|2-warn|3-fatal")
	logFile    *string = flag.String("logfile", "log/proxy.log", "go mysql proxy logfile")
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.SetBlockProfileRate(1)
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

	mysql.SetLogger(log.Logger())

	go func() {
		err := cluster.DisasterControl()
		if err != nil {
			log.Warn(err)
		}
	}()
	go func() {
		// log.info("start checking config file")
		cfg.CheckConfigUpdate(cluster.NotifyChan)
	}()

	sc := make(chan os.Signal, 1)
	Restart := make(chan os.Signal, 1)
	signal.Notify(Restart, syscall.SIGINT)
	signal.Notify(sc, syscall.SIGQUIT)
	// signal.Notify(Restart, syscall.SIGHUP)
	// signal.Notify(sc, syscall.SIGINT)
	// syscall.SIGHUP,
	// syscall.SIGINT,
	//syscall.SIGTERM,

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
		select {
		case sig := <-sc:
			log.Infof("Got signal [%d] to exit.", sig)
			svr.Close()
		case sig := <-Restart:
			log.Infof("Got signal [%d] to Restart.", sig)
			svr.Restart()
		}
	}()

	svr.Serve()
	os.Exit(0)

}

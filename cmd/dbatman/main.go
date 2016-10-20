package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"

	"github.com/bytedance/dbatman/config"
	"github.com/bytedance/dbatman/database/cluster"
	"github.com/bytedance/dbatman/database/mysql"
	"github.com/bytedance/dbatman/proxy"
	"github.com/ngaut/log"
)

var (
	configFile *string = flag.String("config", getCurrentDir()+"/proxy.yml", "go mysql proxy config file")
	logLevel   *int    = flag.Int("loglevel", 0, "0-debug| 1-notice|2-warn|3-fatal")
	logFile    *string = flag.String("logfile", getCurrentDir()+"/proxy.log", "go mysql proxy logfile")
	gcLevel    *string = flag.String("gclevel", "500", "go gc level")
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
func getCurrentDir() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	path1 := substr(path, 0, strings.LastIndex(path, "/"))
	return path1
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.SetBlockProfileRate(1)
	os.Setenv("GOGC", "100")
	log.SetOutputByName(*logFile)
	flag.Parse() //parse tue input argument
	println(*logFile)
	println(*configFile)

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
	signal.Notify(Restart, syscall.SIGUSR1)
	signal.Notify(sc, syscall.SIGQUIT,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM)

	var svr *proxy.Server
	svr, err = proxy.NewServer(cfg)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	//port for go pprof Debug
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

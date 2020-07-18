package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"git.kanosolution.net/kano/appkit"
	"github.com/ariefdarmawan/batcher"
	"github.com/ariefdarmawan/datahub"
	_ "github.com/ariefdarmawan/flexmgo"
	"github.com/ariefdarmawan/kconfigurator"
)

var (
	h         *datahub.Hub
	appConfig = new(kconfigurator.AppConfig)
	config    = flag.String("config", "config/app.yml", "location of config file")
)

func main() {
	flag.Parse()
	e := appkit.ReadConfig(*config, appConfig)
	if e != nil {
		fmt.Println("error:", e.Error())
		os.Exit(1)
	}
	h, e = kconfigurator.MakeHub(appConfig, "default")
	if e != nil {
		fmt.Println("error:", e.Error())
		os.Exit(1)
	}
	defer h.Close()
	h.DeleteQuery(new(batcher.Process), nil)

	mux := http.NewServeMux()
	handlingMux(mux)

	// run service
	csign := make(chan os.Signal)
	go func() {
		hostName := appConfig.Hosts["api"]
		fmt.Printf("Running %v service on %s\n", "api", hostName)
		err := http.ListenAndServe(hostName, mux)
		if err != nil {
			csign <- syscall.SIGINT
		}
	}()

	// grace shutdown
	signal.Notify(csign, os.Interrupt, os.Kill)
	<-csign
	fmt.Printf("Stopping %v service\n", "api")
}

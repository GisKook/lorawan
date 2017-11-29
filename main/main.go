package main

import (
	"fmt"
	"github.com/giskook/lorawan/conf"
	"github.com/giskook/lorawan/http_srv"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cfg, err := conf.ReadConfig("./conf.json")
	if err != nil {
		log.Println(err.Error())
	}

	http_srv := http_srv.NewHttpSrv(cfg)
	http_srv.Start()

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}

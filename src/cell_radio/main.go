package main

import (
	"flag"
	"strings"
	"log"
	"os"
	"cell_radio/receiver"
	"cell_radio/watcher"
	"os/signal"
)

type Conf struct {
	host string
	httpPort string
	httpPath string
	wsPort string
	wsPath string
}

var conf Conf

func init()  {
	flag.StringVar(&conf.host, `host`, `localhost`, `Host to listen`)
	flag.StringVar(&conf.httpPort, `httpPort`, `8080`, `Port to listen HTTP`)
	flag.StringVar(&conf.httpPath, `httpPath`, `ht`, `Path to listen HTTP`)
	flag.StringVar(&conf.wsPort, `wsPort`, `8081`, `Port to serve websockets`)
	flag.StringVar(&conf.wsPath, `wsPath`, `ws`, `Path to serve websockets`)
	flag.Parse()
}

func main()  {
	var (
		logger *log.Logger
		httpReceiver *receiver.Receiver
		wsWatcher *watcher.Watcher
		requestPipe chan interface{}
		shutdownChannel chan os.Signal
	)

	flag.Parse()
	logger = log.New(os.Stderr, `cell_radio server`, log.LstdFlags)
	requestPipe = make(chan interface{}, 100)
	shutdownChannel = make(chan os.Signal, 1)

	httpReceiver = receiver.New(strings.Join([]string{conf.host, conf.httpPort}, `:`), conf.httpPath, &requestPipe, logger)
	wsWatcher = watcher.New(strings.Join([]string{conf.host, conf.wsPort}, `:`), conf.wsPath, &requestPipe, logger)

	httpReceiver.Start()
	wsWatcher.Start()

	defer func() {
		httpReceiver.Shutdown()
		wsWatcher.Shutdown()
	}()

	signal.Notify(shutdownChannel, os.Interrupt)

	<- shutdownChannel
}
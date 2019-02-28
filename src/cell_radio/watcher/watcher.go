package watcher

import (
	"context"
	"net/http"
	"golang.org/x/net/websocket"
	"strings"
	"cell_radio/common"
	"encoding/json"
	"log"
	"sync"
)

type Watcher struct {
	Server *http.Server
	Pipe *chan interface{}
	mx sync.Mutex
	Conns map[*websocket.Conn]struct{}
	logger *log.Logger
}

func New(addr string, path string, pipe *chan interface{}, logger *log.Logger) (w *Watcher) {
	w = &Watcher{}
	w.Server = &http.Server{Addr: addr, Handler: nil}
	w.Pipe = pipe
	w.Conns = make(map[*websocket.Conn]struct{})
	w.logger = logger
	handledPath := strings.Join([]string{``, path, `lastpoints`}, `/`)
	http.Handle(handledPath, websocket.Handler(w.handleLastPointsConnection))
	w.logger.Printf(`Handle WS requests on %v`, handledPath)
	return w
}

func (w *Watcher) Start() {
	go w.Server.ListenAndServe()
	w.logger.Printf(`Started WS watcher server on %v`, w.Server.Addr)
	go w.BroadcastLastPoints()
}

func (w *Watcher) handleLastPointsConnection( conn *websocket.Conn)  {
	w.logger.Println(`Run handleLastPointsConnection...`)
	w.mx.Lock()
	w.Conns[conn] = struct{}{}
	w.mx.Unlock()
	for {
		var msg []byte
		_, err := conn.Read(msg)
		if err != nil {
			w.logger.Printf(`An error occured during connection reading! Error: %v`, err.Error())
			break
		}
	}
	w.logger.Println(`Stop handleLastPointsConnection!`)
}

func (w *Watcher) BroadcastLastPoints()  {
	var (
		respBody []byte
		err error
		respRaw interface{}
		respTyped common.AddManyRequest
		ok bool
	)
	w.logger.Println(`Run BroadcastLastPoints...`)
	for respRaw = range *w.Pipe {
		w.logger.Println(`Received data from pipe!`)
		respTyped, ok = respRaw.(common.AddManyRequest);
		if !ok {
			w.logger.Println(`Wrong reqiest received! Continuing listening...`)
			continue
		}
		respBody, err = json.Marshal(respTyped)
		if err != nil {
			w.logger.Printf(`Could not marshal data! Error: %v`, err.Error())
			continue
		}
		for conn, _ := range w.Conns {
			w.logger.Printf(`Writing data to connection at %v ()...`, conn.Config().Origin.String(), conn.Config().Location.String())
			_, err = conn.Write(respBody)
			if err != nil {
				w.logger.Printf(`Could nor write data to connection! Error: %v`, err.Error())
				w.mx.Lock()
				conn.Close()
				delete(w.Conns, conn)
				w.mx.Unlock()
			}
		}
	}
}

func (w *Watcher) Shutdown() {
	for conn, _ := range w.Conns {
		conn.Close()
		delete(w.Conns, conn)
	}
	w.logger.Println(`Stopping WS server...`)
	w.Server.Shutdown(context.Background())
	w.logger.Println(`WS server stopped!`)
}
package receiver

import (
	"context"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"cell_radio/common"
	"log"
	"strings"
)

type Receiver struct {
	Server *http.Server
	Pipe *chan interface{}
	logger *log.Logger
}

func New(addr string, path string, pipe *chan interface{}, logger *log.Logger) (r *Receiver) {
	r = &Receiver{}
	r.Server = &http.Server{Addr: addr, Handler: nil}
	r.Pipe = pipe
	r.logger = logger
	handledPath := strings.Join([]string{``, path, `addmany`}, `/`)
	http.HandleFunc(handledPath, r.HandleAddMany)
	r.logger.Printf(`Handle HTTP requests on %v`, handledPath)
	http.HandleFunc(`/`, func(w http.ResponseWriter, req *http.Request)  {r.logger.Printf(`Accepted on %v`, req.RequestURI)})
	r.logger.Println(`Handle HTTP requests on /`)
	return r
}

func (r *Receiver) Start() {
	r.logger.Printf(`Started HTTP receiver server on %v`, r.Server.Addr)
	go r.Server.ListenAndServe()
}

func (r *Receiver) HandleAddMany(w http.ResponseWriter, req *http.Request)  {
	var (
		reqBody, respBody []byte
		err error
		data common.AddManyRequest
	)
	r.logger.Println(`Starting HandleAddMany...`)
	defer req.Body.Close()
	reqBody, err = ioutil.ReadAll(req.Body)
	if err != nil {
		r.logger.Printf(`Could not read request body! error: %v`, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		r.logger.Printf(`Could not parse request body! error: %v`, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	*r.Pipe <- data
	respBody, err = json.Marshal(r.makeResponse(data))
	if err != nil {
		r.logger.Printf(`Could not marshal response body! error: %v`, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(respBody)
}

func (r *Receiver) makeResponse(reqData common.AddManyRequest) (respData common.AddManyResponse) {
	respData.UsrpCfg = common.UsrpCfg{
		Config: reqData.Config,
		Watcher: 0,
		ScanPsc: `0`,
		ScanTimeouts: `0`,
		ScanUlsc: `0`,
		Mode: 0,
	}
	respData.Command = 0
	return respData
}

func (r *Receiver) Shutdown() {
	r.logger.Println(`Stopping HTTP server...`)
	r.Server.Shutdown(context.Background())
	r.logger.Println(` HTTP server stopped!`)
}


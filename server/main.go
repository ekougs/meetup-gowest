package main

import (
	"net/http"
	"log"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Meteo struct {}

type NoArgs struct {}

type Reply struct {
	Result int
}

func (o *Meteo) GetRainInfo(r *http.Request, args *NoArgs, reply *Reply) error {
	reply.Result = 5
	return nil
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterService(new(Meteo), "")
	http.Handle("/rpc", rpcServer)

	log.Println("Server started")
	if e := http.ListenAndServe(":2015", nil); e != nil {
		log.Fatal(e)
	}
}

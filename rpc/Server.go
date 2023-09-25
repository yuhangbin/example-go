package rpc

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const serverAddress = "127.0.0.1"

func server() {

	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", serverAddress+":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}

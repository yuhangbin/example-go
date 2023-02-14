package main

import (
	"example/rpc/core"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const serverAddress = "127.0.0.1"

func main() {
	arith := new(core.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", serverAddress+":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}

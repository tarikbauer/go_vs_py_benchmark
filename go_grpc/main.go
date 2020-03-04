package main

import (
	"log"
	"net"
	"os"

	"github.com/tarikbauer/go_vs_py_benchmark/go_grpc/api"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("GO_GRPC_PORT")
	listener, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	server := api.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterTimeEvaluationServer(grpcServer, &server)
	api.LogInfo("Server Listening on 127.0.0.1:" + port)
	log.Fatal(grpcServer.Serve(listener))
}
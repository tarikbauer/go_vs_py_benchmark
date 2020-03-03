package main

import (
	"go_grpc/api"
	"log"
	"net"
	"os"

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
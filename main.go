package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/tarikbauer/go_vs_py_benchmark/client"
	"google.golang.org/grpc"
)

const ITERATIONS int = 20
var VALUES = "1,2,1,1,5,1,2,3"

func getUrl(port string) string {
	return "http://127.0.0.1:" + port + "/api?t=" + VALUES
}

func getGRPCConn() *grpc.ClientConn {
	conn, err := grpc.Dial(":" + os.Getenv("GO_GRPC_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatal("GRPC connection failed!")
	}
	return conn
}

func getRESTConn() *http.Client {
	return &http.Client{}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed loading `.env` file!")
	}

	client.BenchMark{
		Value:      getUrl(os.Getenv("SANIC_PORT")),
		Iterations: ITERATIONS,
		Name:       "Sanic",
		Channel:    make(chan float64, ITERATIONS),
		WG:         sync.WaitGroup{},
		Client:     client.RESTConn{Conn: getRESTConn()},
	}.Run()
	client.BenchMark{
		Value:      getUrl(os.Getenv("GO_MUX_PORT")),
		Iterations: ITERATIONS,
		Name:       "Go Mux",
		Channel:    make(chan float64, ITERATIONS),
		WG:         sync.WaitGroup{},
		Client:     client.RESTConn{Conn: getRESTConn()},
	}.Run()
	client.BenchMark{
		Value: VALUES, Iterations: ITERATIONS, Name: "Go GRPC",
		Channel: make(chan float64, ITERATIONS),
		WG:      sync.WaitGroup{},
		Client:  client.GRPCConn{Conn: getGRPCConn()},
	}.Run()
}

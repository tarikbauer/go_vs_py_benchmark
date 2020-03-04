package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const ITERATIONS int = 50
var VALUES = "1,1,1,1"

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

	pythonSanicBenchmark := helper.BenchMark{
		getUrl(os.Getenv("SANIC_PORT")),
		ITERATIONS,
		"Sanic",
		make(chan float64, ITERATIONS),
		sync.WaitGroup{},
		helper.RESTConn{getRESTConn()},
	}
	goMuxBenchmark := helper.BenchMark{
		getUrl(os.Getenv("GO_MUX_PORT")),
		ITERATIONS,
		"Go Mux",
		make(chan float64, ITERATIONS),
		sync.WaitGroup{},
		helper.RESTConn{getRESTConn()},
	}
	goGRPCBenchmark := helper.BenchMark{
		VALUES, "Go GRPC",
		ITERATIONS,
		make(chan float64, ITERATIONS),
		sync.WaitGroup{},
		helper.GRPCConn{getGRPCConn()},
	}
	pythonSanicBenchmark.run()
	pythonSanicBenchmark.displayResult()
	goMuxBenchmark.run()
	goMuxBenchmark.displayResult()
	goGRPCBenchmark.run()
	goGRPCBenchmark.displayResult()
}

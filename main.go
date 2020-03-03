package main

import (
	"context"
	"fmt"
	"go_grpc/api"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const ITERATIONS int = 20
var VALUES = "1,1,1,1"

type InvalidResponse struct{}

func (InvalidResponse) Error() string {
	return "Invalid Response Status Code!"
}

type Client interface {
	fetch(input string) (float64, error)
	io.Closer
}

type GRPCConn struct {
	conn *grpc.ClientConn
}

type RESTConn struct {
	conn *http.Client
}

func (c GRPCConn) Close() error {
	err := c.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c GRPCConn) fetch(input string) (float64, error) {
	var values []int64
	client := api.NewTimeEvaluationClient(c.conn)
	for _, value := range strings.Split(input, ",") {
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0, err
		}
		values = append(values, value)
	}
	t := time.Now()
	_, err := client.Evaluate(context.Background(), &api.TimeRequest{Values: values})
	elapsed := time.Since(t)
	return float64(elapsed), err
}

func (c RESTConn) Close() error {
	c.conn.CloseIdleConnections()
	return nil
}

func (c RESTConn) fetch(input string) (float64, error) {
	t := time.Now()
	response, err := c.conn.Get(input)
	elapsed := time.Since(t)
	if err != nil{
		return 0, err
	} else if response.StatusCode >= 300{
		return 0, InvalidResponse{}
	}
	return float64(elapsed), nil
}

type BenchMark struct {
	value string
	name string
	c chan float64
	wg sync.WaitGroup
	client Client
}

func (b *BenchMark) evaluate(input string) {
	defer b.wg.Done()
	elapsed, err := b.client.fetch(input)
	if err != nil{
		log.Fatal(b.name + ": ", err)
	}
	b.c <- elapsed
}

func (b *BenchMark) run() {
	defer b.closeAll()
	for i := 0; i < ITERATIONS; i++ {
		b.wg.Add(1)
		go b.evaluate(b.value)
	}
	b.wg.Wait()
}

func (b BenchMark) displayResult() {
	for value := range b.c {
		fmt.Print(b.name + ": ")
		fmt.Println(value)
	}
}

func (b *BenchMark) closeAll() {
	close(b.c)
	err := b.client.Close()
	if err != nil {
		log.Fatal("Error while closing connection!")
	}
}

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

	pythonSanicBenchmark := BenchMark{
		getUrl(os.Getenv("SANIC_PORT")),
		"Sanic",
		make(chan float64, ITERATIONS),
		sync.WaitGroup{},
		RESTConn{getRESTConn()},
	}
	goMuxBenchmark := BenchMark{
		getUrl(os.Getenv("GO_MUX_PORT")),
		"Go Mux",
		make(chan float64, ITERATIONS),
		sync.WaitGroup{},
		RESTConn{getRESTConn()},
	}
	goGRPCBenchmark := BenchMark{
		VALUES, "Go GRPC",
		make(chan float64, ITERATIONS),
		sync.WaitGroup{},
		GRPCConn{getGRPCConn()},
	}
	pythonSanicBenchmark.run()
	pythonSanicBenchmark.displayResult()
	goMuxBenchmark.run()
	goMuxBenchmark.displayResult()
	goGRPCBenchmark.run()
	goGRPCBenchmark.displayResult()
}

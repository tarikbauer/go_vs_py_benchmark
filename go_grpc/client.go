package main

import (
	"context"
	"fmt"
	"go_grpc/api"
	"google.golang.org/grpc"
	"log"
	"os"
)

func closeConn(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Fatal("Error while closing GRPC connection!")
	}
}

func main() {
	port := os.Getenv("PORT")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":" + port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Connection Failed!")
	}
	defer closeConn(conn)
	client := api.NewTimeEvaluationClient(conn)
	response, err := client.Evaluate(context.Background(), &api.TimeRequest{Values: []int64{1, 5, 3, 2}})
	if err != nil {
		log.Fatal("Request Error!")
	}
	fmt.Println(response.Response)
}

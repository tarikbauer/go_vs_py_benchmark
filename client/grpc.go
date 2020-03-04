package client

import (
	"context"
	"google.golang.org/grpc"
	"time"

	"github.com/tarikbauer/go_vs_py_benchmark/go_grpc/api"
)

type GRPCConn struct {
	Conn *grpc.ClientConn
}

func (c GRPCConn) Close() error {
	err := c.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c GRPCConn) fetch(input string) (float64, error) {
	client := api.NewTimeEvaluationClient(c.Conn)
	values, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	t := time.Now()
	_, err = client.Evaluate(context.Background(), &api.TimeRequest{Values: values})
	elapsed := time.Since(t)
	return float64(elapsed), err
}

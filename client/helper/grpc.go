package helper

import (
	"context"
	"google.golang.org/grpc"
	"strconv"
	"strings"
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
	var values []int64
	client := api.NewTimeEvaluationClient(c.Conn)
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

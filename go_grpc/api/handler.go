package api

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"
)

type Server struct {}

func Sleep(t int64, c chan <- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(t) * time.Second)
	c <- t
}

func (s *Server) Evaluate(_ context.Context, r *TimeRequest) (*EvaluationResponse, error) {
	LogInfo("Request received: " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(r.Values)), ","), "[]"))
	var wg sync.WaitGroup
	c := make(chan int64, len(r.Values))
	for _, value := range r.Values {
		wg.Add(1)
		go Sleep(value, c, &wg)
	}
	wg.Wait()
	close(c)
	var response int64 = 0
	for value := range c{
		response += value
	}
	return &EvaluationResponse{Response: response}, nil
}

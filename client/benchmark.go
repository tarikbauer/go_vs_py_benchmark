package client

import (
	"fmt"
	"io"
	"log"
	"sync"
)

type Client interface {
	fetch(input string) (float64, error)
	io.Closer
}

type BenchMark struct {
	Value      string
	Iterations int
	Name       string
	Channel    chan float64
	WG         sync.WaitGroup
	Client     Client
}

func (b *BenchMark) Evaluate(input string) {
	defer b.WG.Done()
	elapsed, err := b.Client.fetch(input)
	if err != nil{
		log.Fatal(b.Name + ": ", err)
	}
	b.Channel <- elapsed
}

func (b *BenchMark) Run() {
	defer b.CloseAll()
	for i := 0; i < b.Iterations; i++ {
		b.WG.Add(1)
		go b.Evaluate(b.Value)
	}
	b.WG.Wait()
}

func (b BenchMark) DisplayResult() {
	for value := range b.Channel {
		fmt.Print(b.Name + ": ")
		fmt.Println(value)
	}
}

func (b *BenchMark) CloseAll() {
	close(b.Channel)
	err := b.Client.Close()
	if err != nil {
		log.Fatal("Error while closing connection!")
	}
}

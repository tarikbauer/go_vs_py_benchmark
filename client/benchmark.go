package client

import (
	"io"
	"log"
	"math"
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
	Results    []float64
}

func (b *BenchMark) Evaluate(input string) {
	defer b.WG.Done()
	elapsed, err := b.Client.fetch(input)
	if err != nil{
		log.Fatal(b.Name + ": ", err)
	}
	b.Channel <- elapsed / math.Pow(10, 9)
}

func (b *BenchMark) Work() {
	defer b.CloseAll()
	for i := 0; i < b.Iterations; i++ {
		b.WG.Add(1)
		go b.Evaluate(b.Value)
	}
	b.WG.Wait()
}

func (b *BenchMark) DisplayResults() {
	for _, value := range b.Results {
		printWarning(b.Name + ": ", value)
	}
}

func (b *BenchMark) GetDelay() float64 {
	var delay float64
	for _, value := range b.Results {
		delay += value
	}
	return delay
}

func (b *BenchMark) DisplayTotalDelay() {
	printInfo(b.Name + " total delay: ", b.GetDelay())
}

func (b *BenchMark) setResults() {
	for value := range b.Channel {
		b.Results = append(b.Results, value)
	}
}

func (b *BenchMark) CloseAll() {
	defer b.setResults()
	close(b.Channel)
	err := b.Client.Close()
	if err != nil {
		log.Fatal("Error while closing connection!")
	}
}

func (b BenchMark) Run() {
	b.Work()
	b.DisplayResults()
	b.DisplayTotalDelay()
}

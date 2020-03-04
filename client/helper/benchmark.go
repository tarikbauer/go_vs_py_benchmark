package helper

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
	value string
	iterations int
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
	for i := 0; i < b.iterations; i++ {
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

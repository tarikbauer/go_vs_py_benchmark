package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	InfoColor   = "\033[0;36m%s\033[0m"
	WarningColor = "\033[0;33m%s\033[0m"
)

func GetIsoFormat() string {
	return time.Now().Format(time.RFC3339)
}

func LogInfo(message interface{}) {
	fmt.Printf(InfoColor, "[" + GetIsoFormat() + "] - [INFO] ")
	fmt.Println(message)
}

func LogWarning(message interface{}) {
	fmt.Printf(WarningColor, "[" + GetIsoFormat() + "] - [WARNING] ")
	fmt.Println(message)
}

func Sleep(t int, c chan <- int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(t) * time.Millisecond)
	c <- t
}

package client

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	InfoColor   = "\033[0;36m%s\033[0m"
	WarningColor = "\033[0;33m%s\033[0m"
)

func printInfo(message string, value float64) {
	fmt.Print(message)
	fmt.Printf(InfoColor, strconv.FormatFloat(value, 'e', -1, 64) + "s\n")
}

func printWarning(message string, value float64) {
	fmt.Print(message)
	fmt.Printf(WarningColor, strconv.FormatFloat(value, 'e', -1, 64) + "s\n")
}

func parseInput(input string) ([]int64, error) {
	var values []int64
	for i, value := range strings.Split(input, ",") {
		if i == 0 {
			value = string(value[len(value) - 1])
		}
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		values = append(values, value)
	}
	return values, nil
}

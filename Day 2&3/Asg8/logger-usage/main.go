package main

import (
	"fmt"
	"logger-usage/logger"
)

func main() {
	logStore := logger.New()
	fmt.Println(logStore)
}
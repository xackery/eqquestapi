package main

import (
	"fmt"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed:", err.Error())
		os.Exit(1)
	}
}

func run() error {
	return nil
}

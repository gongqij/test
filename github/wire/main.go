package main

import (
	"fmt"
	"os"
)

func main() {
	e, err := InitializeEvent("我是龚琪，我是一个大帅比")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}

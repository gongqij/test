package main

import (
	"fmt"
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := Get("http://127.0.0.1:8080/a")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Got resp:", string(resp))
}

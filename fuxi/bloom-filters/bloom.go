package main

import (
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	filter := bloom.New(1000, 4)
	filter.Add([]byte("Love"))
	if filter.Test([]byte("Love")) {
		fmt.Println("love is exist")
	}
}

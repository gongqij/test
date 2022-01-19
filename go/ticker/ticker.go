package main

import (
	"fmt"
	"time"
)

func main() {
	flag := true
	tic := time.NewTicker(time.Second * 5)
	defer tic.Stop()
	for {
		select {
		case <-tic.C:
			if flag {
				time.Sleep(time.Second * 3)
				flag = false
			}
			fmt.Println(time.Now())
		}
	}
}

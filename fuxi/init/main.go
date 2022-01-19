package main

import (
	"fmt"
	"go-study/fuxi/init/a"
	"go-study/fuxi/init/b"
)

func main() {
	fmt.Println("这里是 services run")
	a.TestA()
	b.TestB()

}
func init() {
	fmt.Println("这里是 services里面的init")
}

func init() {
	fmt.Println("xxxx")
}

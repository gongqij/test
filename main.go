package main

import "fmt"

func main() {
	var buf []byte
	add(&buf, "foo")
	add(&buf, []byte("bar"))
	fmt.Println(string(buf))
}

func add[S ~string | ~[]byte](buf *[]byte, s S) {
	*buf = append(*buf, []byte(s)...)
}

package main

import (
	"constraints"
	"fmt"
)

func main() {
	fmt.Println(min("yyy", "xxxx"))
	var buf []byte
	add(&buf, "foo")
	add(&buf, []byte("bar"))
	fmt.Println(string(buf))
}

func min[P constraints.Ordered](x, y P) P {
	if x < y {
		return x
	} else {
		return y
	}
}

func add[S ~string | ~[]byte](buf *[]byte, s S) {
	*buf = append(*buf, []byte(s)...)
}

package main

import (
	"bou.ke/monkey"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMonkey(t *testing.T) {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?") // what the *bleep*?
	fmt.Println("what the hell?")
}

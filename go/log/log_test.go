package log

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println("TestA running")
}

func TestB(t *testing.T) {
	fmt.Println("TestB running")
}

func TestMain(m *testing.M) {
	fmt.Println("Do stuff BEFORE the tests!")
	_ = m.Run()
	fmt.Println("Do stuff AFTER the tests!")

}

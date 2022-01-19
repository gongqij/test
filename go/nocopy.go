package nocopy

import (
	"fmt"
)

// noCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
type noCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type Test struct {
	noCopy noCopy
	ID     int
	Name   string
}

func HelloWorld() {
	fmt.Println("hello world")
}

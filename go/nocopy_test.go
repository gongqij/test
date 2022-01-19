package nocopy

import (
	"fmt"
	"testing"
)

func TestNoCopy(t *testing.T) {
	no := Test{
		ID:   1,
		Name: "aasd",
	}

	fmt.Println(no)
}

func BenchmarkNoCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld()
	}
}

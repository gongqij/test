package pool

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	sli := make([]int, 0, 10)
	sli2 := sli[:5]
	fmt.Println(sli2)
	for i := 1; i <= 100; i++ {
		pool(i)
	}
}

func BenchmarkPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool(i)
	}
}

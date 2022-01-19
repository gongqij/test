package array

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	linkList := &List{}
	linkList.init([]int{1, 2, 5, 3, 4})
	linkList.print()
}

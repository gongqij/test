package main

import (
	"fmt"
	"go-study/fuxi/sort"
)

func main() {
	sli := []int{5, 2, 1, 4, 3}
	sort.HeapSort(sli)
	fmt.Println(sli)
}

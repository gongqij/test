package sort

//HeapSort
//堆排序的平均时间复杂度为 Ο(nlogn)
func HeapSort(arr []int) {
	arrLen := len(arr)
	// 建立大顶堆
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
	for i := arrLen - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		arrLen -= 1
		//调整顶点就可以使堆继续保持大顶堆状态
		heapify(arr, 0, arrLen)
	}
}

// 堆调整
func heapify(arr []int, i, arrLen int) {
	//找父节点、左节点、右节点中最大的那个
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	//如果发现左节点或右节点比父节点大，交换父子节点，然后递归调整该子节点的子节点
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, largest, arrLen)
	}
}

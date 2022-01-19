package sort

//QuickSort
//快速排序的最坏运行情况是 O(n²)，比如说顺序数列的快排。
//但它的平摊期望时间是 O(nlogn)，且 O(nlogn) 记号中隐含的常数因子很小，比复杂度稳定等于 O(nlogn) 的归并排序要小很多。
//所以，对绝大多数顺序性较弱的随机数列而言，快速排序总是优于归并排序。
func QuickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			_swap(arr, i, index)
			index += 1
		}
	}
	_swap(arr, pivot, index-1)
	return index - 1
}

func _swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

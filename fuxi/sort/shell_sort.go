package sort

//希尔排序的时间复杂度与增量(即，步长gap)的选取有关。
//例如，当增量为1时，希尔排序退化成了直接插入排序，此时的时间复杂度为O(N²)，
//而Hibbard增量的希尔排序的时间复杂度为O(N3/2)。
//希尔排序是不稳定的算法
//在此我们选择增量gap=length/2，缩小增量继续以gap = gap/2的方式，
//这种增量选择我们可以用一个序列来表示，{n/2,(n/2)/2...1}，称为增量序列。
//希尔排序的增量序列的选择与证明是个数学难题，我们选择的这个增量序列是比较常用的，也是希尔建议的增量，称为希尔增量，
//但其实这个增量序列不是最优的。此处我们做示例使用希尔增量
//ShellSort
func ShellSort(arr []int) []int {
	length := len(arr)
	gap := length / 2 //分为gap个组
	for gap > 0 {
		//将下标为i=gap、gap+1、gap+2...length-1的值arr[i]依次向前插入(直接插入排序)
		//保证各个组组内有序。例如有6个数字，分为3组下标依次为(0、3)，(1、4)，(2、5)，
		//循环结束后保证arr[0]>=arr[3],arr[1]>=arr[4]，arr[2]>=arr[5]。
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 2
	}
	return arr
}

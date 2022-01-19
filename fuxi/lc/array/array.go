package array

/*
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]
*/
//用map的索引，省去内层循环，使得时间复杂度小于 O(n2)
func twoSum(nums []int, target int) []int {
	var result []int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			tempSlice := []int{m[another], i}
			result = tempSlice
		}
		m[nums[i]] = i
	}
	return result
}

func twoSum2(nums []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

//删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[index+1] = nums[i]
			index++
		}
	}

	return index + 1
}

//移除元素
func removeElement(nums []int, val int) ([]int, int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return nums[:index], index
}

func removeElement2(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for {
		if i >= j {
			break
		}
		for i <= len(nums)-1 && nums[i] != val {
			i++
		}
		for j >= 0 && nums[j] == val {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--

	}
	return i
}

//最长连续子序列
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	max := 0
	for k := range m {
		//如果存在前驱跳过
		if m[k-1] {
			continue
		}
		current := 1
		next := k + 1
		for m[next] {
			next++
			current++
		}
		if current > max {
			max = current
		}
	}
	return max
}

//最大连续子序和
func maxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		temp := 0
		for j := i; j < len(nums); j++ {
			temp = temp + nums[j]
			if max < temp {
				max = temp
			}
		}
	}
	return max
}

//最大连续子序和
//动态规划
func maxSubArray2(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

// 搜索插入位置
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

// 加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	return append([]int{1}, digits...)
}

//合并两个有序数组
//正序遍历
func merge1(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, tail := m-1, n-1, m+n-1
	for {
		if p2 == -1 {
			break
		}
		if p1 == -1 {
			for ; p2 >= 0; p2-- {
				nums1[tail] = nums2[p2]
				tail--
			}
			break
		}
		if nums1[p1] < nums2[p2] {
			nums1[tail] = nums2[p2]
			p2--
			tail--
		} else {
			nums1[tail] = nums1[p1]
			p1--
			tail--
		}
	}
}

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNumber(nums []int) int {
	singleNumber := 0
	for i := 0; i < len(nums); i++ {
		flag := true
		for j := 0; j < len(nums); j++ {
			if i == j {
				if i == len(nums)-1 {
					singleNumber = nums[i]
					flag = false
				}
				continue
			}
			if nums[i] == nums[j] {
				flag = false
				break
			}
		}
		if flag {
			singleNumber = nums[i]
			break
		}
	}
	return singleNumber
}

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNumber2(nums []int) int {
	singleNumber := 0
	for i := 0; i < len(nums); i++ {
		singleNumber ^= nums[i]
	}
	return singleNumber
}

//螺旋矩阵
func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total          = rows * columns
		order          = make([]int, total)
		row, column    = 0, 0
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}

func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

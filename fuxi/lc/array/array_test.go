package array

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	result := twoSum2([]int{11, 11, 3, 9, 7, 7, 11}, 10)
	fmt.Println(result)
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{}
	len := removeDuplicates(nums)
	for i := 0; i < len; i++ {
		fmt.Println(nums[i])
	}
}

func TestRemoveElement(t *testing.T) {
	//nums := [0,1,2,2,3,0,4,2]
	len := removeElement2([]int{3, 2, 2, 3}, 3)
	fmt.Println(len)
}

func TestLongestConsecutive(t *testing.T) {
	sli := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	num := longestConsecutive(sli)
	fmt.Println(num)
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
}

func TestSingleNumber(t *testing.T) {
	fmt.Println(spiralOrder1([][]int{{3}, {2}}))
}

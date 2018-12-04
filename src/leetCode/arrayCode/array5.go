/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
*/
package main

import "sort"

func singleNumber(nums []int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		k ^= nums[i]
	}
	return k
}

//func singleNumber(nums []int) int {
//	sort.Ints(nums)
//	for i := 0; i < len(nums); i += 2 {
//		if i+1 >= len(nums) {
//			return nums[i]
//		}
//		if nums[i] != nums[i+1] {
//			return nums[i]
//		}
//	}
//	return -1
//}

func main() {
	arr := []int{4, 1, 2, 1, 2}
	sort.Ints(arr)
	singleNumber(arr)
}

/**
给定一个整数数组，判断是否存在重复元素。

如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
示例 1:

输入: [1,2,3,1]
输出: true
*/
package main

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		var target = nums[i]
		for j := i + 1; j < len(nums); j++ {
			if target == nums[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	temp := []int{1, 2, 3, 4}
	containsDuplicate(temp)
}

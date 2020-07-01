package leetcode

import "sort"

//169. 多数元素

/*
	给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

	你可以假设数组是非空的，并且给定的数组总是存在多数元素。



	示例 1:

	输入: [3,2,3]
	输出: 3
	示例 2:

	输入: [2,2,1,1,1,2,2]
	输出: 2
*/

//1.暴力破解法，利用双重for循环统计

//2.利用map统计元素出现的次数
func majorityElement(nums []int) int {
	n := len(nums)

	if n == 0 {
		return -1
	}

	if n < 2 {
		return nums[0]
	}

	hash := make(map[int]int)

	for i := 0; i < n; i++ {
		hash[nums[i]]++
		count, _ := hash[nums[i]]
		if count > n/2 {
			return nums[i]
		}
	}

	return -1
}

//3.先对切片进行排序
func majorityElement1(nums []int) int {

	n := len(nums)

	if n == 0 {
		return -1
	}

	if n < 2 {
		return nums[0]
	}

	sort.Ints(nums)

	return nums[n/2]
}

//4.利用大于 ⌊ n/2 ⌋ 的特性
func majorityElement2(nums []int) int {

	n := len(nums)

	if n == 0 {
		return -1
	}

	if n < 2 {
		return nums[0]
	}

	value, count := 0, 0

	for i := 0; i < n; i++ {
		if count == 0 {
			value, count = nums[i], 1
		} else if value == nums[i] {
			count++
		} else {
			count--
		}
	}
	return value
}

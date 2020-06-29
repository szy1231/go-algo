/*
@Time : 2020/4/25 9:35
@Author : pasiyu
@File : binarysearch.go
*/
package binarysearch


//递归实现
func BinarySearch(a []int, value int) int {
	if len(a) <= 0 {
		return -1
	}
	return binarySearch(a, 0, len(a)-1, value)
}

func binarySearch(a []int, low, high, value int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)/2

	if a[mid] == value {
		return mid
	} else if a[mid] > value {
		return binarySearch(a, low, mid-1, value)
	} else {
		return binarySearch(a, mid+1, high, value)
	}
}

//循环实现
func BinarySearchCycle(a []int, v int) int {
	n := len(a)
	if n <= 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)/2

		if a[mid] == v {
			return mid
		} else if a[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

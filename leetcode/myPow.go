package leetcode

//50. Pow(x, n)：
/*
	实现 pow(x, n) ，即计算 x 的 n 次幂函数。

	示例 1:

	输入: 2.00000, 10
	输出: 1024.00000
	示例 2:

	输入: 2.10000, 3
	输出: 9.26100
	示例 3:

	输入: 2.00000, -2
	输出: 0.25000
	解释: 2-2 = 1/22 = 1/4 = 0.25
	说明:

	-100.0 < x < 100.0
	n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/

//解题思路：
//分治思想，利用x^y = x^2/y * x^2/y
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var pow float64 = 1
	for n > 0 {
		//如果n为奇数，则先乘以x
		if n&1 == 1 {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}

//递归实现：
func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	return quickMul(x, n)
}

func quickMul(x float64, n int) float64 {
	if n <= 1 {
		return x
	}
	pow := 1.0
	if n&1 == 1 {
		pow *= x
	}
	x *= x
	n >>= 1
	return quickMul(x, n) * pow
}

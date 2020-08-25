package leetcode

import "strings"

/**

459. 重复的子字符串
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例 1:

输入: "abab"

输出: True

解释: 可由子字符串 "ab" 重复两次构成。
示例 2:

输入: "aba"

输出: False
示例 3:

输入: "abcabcabcabc"

输出: True

解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)

*/

//解法1：新建一个字符串s+s，去掉开头结尾，如果剩下的还包含s。则为true
func repeatedSubstringPattern1(s string) bool {

	str1 := s + s
	str2 := str1[1 : len(str1)-1]

	if strings.Contains(str2, s) {
		return true
	}
	return false

}

//解法2
/**
首先要知道：子字符串的开头和结尾，肯定与传入s的开头和结尾字符相同
循环找到与字符串末尾相同的字符，然后判断其之后字符是否是之前字符的复制版
*/
func repeatedSubstringPattern2(s string) bool {
	n := len(s)
	if n <= 1 {
		return false
	}

	r := 0
	//如果一半都没找到子字符串，则肯定没有。返回false
	for r < n/2 {
		//相同，之前字符串则可能为子字符串。否则继续向后查找 r++
		if s[r] == s[n-1] {
			//传入其 前面字符串和后面字符串。
			if isRepeated(s[0:r+1], s[r+1:]) {
				return true
			}
		}
		r++
	}
	return false

}

func isRepeated(s1, s2 string) bool {
	n := len(s1)
	//如果s2不是s1的倍数，则肯定不是
	if len(s2)%n != 0 {
		return false
	}

	par := 0
	//检查s2是否是s1的复制版
	for par < len(s2) {
		if s1 != s2[par:par+n] {
			return false
		}
		par += n
	}

	return true

}

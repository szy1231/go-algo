package leetcode

//242. 有效的字母异位词
/*给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
*/

//利用map统计字符出现的次数
func isAnagram(s string, t string) bool {
	hash := map[rune]int{}

	for _, v := range s {
		hash[v]++
	}
	for _, v := range t {
		vs, ok := hash[v]
		if !ok {
			return false
		}
		if vs == 1 {
			delete(hash, v)
		} else {
			hash[v]--
		}
	}
	return len(hash) == 0

}

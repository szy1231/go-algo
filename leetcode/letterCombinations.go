package leetcode

/*
	给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

*/

//思路：循环遍历，依次添加符合的元素组合
func letterCombinations(digits string) []string {
	//如果是空，则直接返回
	if len(digits) == 0 {
		return nil
	}

	//创建存储结果的切片
	res := make([]string, 0)
	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	//将第一个元素添加到切片中，方便下面的操作
	for _, v := range m[string(digits[0])] {
		res = append(res, string(v))
	}
	//如果只有一个元素，则返回结果
	if len(digits) == 1 {
		return res
	}

	//循环，依次添加
	for i := 1; i < len(digits); i++ {
		res = combination(res, m[string(digits[i])])
	}

	return res

}

func combination(arr []string, s string) (res []string) {
	for _, v := range s {
		for i := 0; i < len(arr); i++ {
			res = append(res, arr[i]+string(v))
		}
	}
	return
}

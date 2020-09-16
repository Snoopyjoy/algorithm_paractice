package leetcode

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	smap := make(map[rune]int, len(s))

	for _, sv := range s{
		smap[sv]++
	}
	for _, tv := range t{
		smap[tv]--
	}

	for _,v := range smap{
		if v != 0 {
			return false
		}
	}

	return true
}

// @lc code=end'

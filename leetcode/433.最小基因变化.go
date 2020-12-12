/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
var min int = -1
var cs []rune = []rune{'A', 'C', 'G', 'T'}

func minMutation(start string, end string, bank []string) int {
	min = -1
	bankMap := make(map[string]struct{}, len(bank))
	for _, v := range bank {
		bankMap[v] = struct{}{}
	}

	if _, ok := bankMap[end]; !ok {
		return -1
	}
	visited := map[string]struct{}{}
	visited[start] = struct{}{}

	tracking(start, end, bankMap, visited, 0)
	return min
}

func tracking(start string, end string, bankMap map[string]struct{}, visited map[string]struct{}, t int) {
	if start == end {
		if min == -1 || t < min {
			min = t
		}
		return
	}
	for i, v := range start {
		for _, r := range cs {
			if v == r {
				continue
			}
			s := []rune(start)
			s[i] = r
			ss := string(s)
			if _, ok := bankMap[ss]; !ok {
				continue
			}
			if _, ok := visited[ss]; ok {
				continue
			}
			visited[ss] = struct{}{}
			tracking(ss, end, bankMap, visited, t+1)
		}
	}
}

// @lc code=end


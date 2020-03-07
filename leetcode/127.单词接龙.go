package leetcode

/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		wordMap[w] = true
	}
	if !wordMap[endWord] {
		return 0
	}
	beginMap := make(map[string]bool)
	endMap := make(map[string]bool)
	l := 1
	visitedMap := make(map[string]bool)

	beginMap[beginWord] = true
	endMap[endWord] = true
	for len(beginMap) > 0 && len(endMap) > 0 {
		if len(beginMap) > len(endMap) {
			beginMap, endMap = endMap, beginMap
		}
		temp := map[string]bool{}
		for word := range beginMap {
			wordArr := []rune(word)
			for i, r := range wordArr {
				for c := 'a'; c <= 'z'; c++ {
					wordArr[i] = c
					wordTemp := string(wordArr)
					if endMap[wordTemp] {
						return l + 1
					}
					if !visitedMap[wordTemp] && wordMap[wordTemp] {
						temp[wordTemp] = true
						visitedMap[wordTemp] = true
					}
					wordArr[i] = r
				}
			}
		}

		beginMap = temp
		l++
	}
	return 0
}

// @lc code=end

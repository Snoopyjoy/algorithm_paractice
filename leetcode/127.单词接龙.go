package leetcode_127

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
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	beginMap := make(map[string]bool, len(wordList))
	endMap := make(map[string]bool, len(wordList))
	visitedMap := make(map[string]bool, len(wordList))

	beginMap[beginWord] = true
	endMap[endWord] = true
	visitedMap[beginWord] = true
	visitedMap[endWord] = true

	step := 1
	for len(beginMap) > 0 && len(endMap) > 0 {
		if len(beginMap) > len(endMap) {
			beginMap, endMap = endMap, beginMap
		}
		tempMap := make(map[string]bool, len(beginMap))
		for w := range beginMap {
			wArr := []rune(w)

			for i, r := range wArr {
				for b := 'a'; b <= 'z'; b++ {
					wArr[i] = b
					tempWord := string(wArr)

					if !wordMap[tempWord] {
						wArr[i] = r
						continue
					}
					if endMap[tempWord] {
						return step + 1
					}
					if !visitedMap[tempWord] {
						tempMap[tempWord] = true
						visitedMap[tempWord] = true
					}
					wArr[i] = r
				}
			}
		}
		beginMap = tempMap
		step++
	}
	return 0
}

// @lc code=end

// 单向bfs
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		wordMap[w] = true
	}
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	beginMap := make(map[string]bool, len(wordList))
	visitedMap := make(map[string]bool, len(wordList))

	beginMap[beginWord] = true
	visitedMap[beginWord] = true
	visitedMap[endWord] = true

	step := 1
	for len(beginMap) > 0 {
		tempMap := make(map[string]bool, len(beginMap))
		for w := range beginMap {
			wArr := []rune(w)
			for i, r := range wArr {
				for b := 'a'; b <= 'z'; b++ {
					wArr[i] = b
					tempWord := string(wArr)

					if !wordMap[tempWord] {
						wArr[i] = r
						continue
					}
					if tempWord == endWord {
						return step + 1
					}
					if !visitedMap[tempWord] {
						tempMap[tempWord] = true
						visitedMap[tempWord] = true
					}
					wArr[i] = r
				}
			}
		}
		beginMap = tempMap
		step++
	}
	return 0
}

package leetcode_126

/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 */

// @lc code=start
type layerItem struct {
	result      [][]string
	visitedMaps []map[string]bool
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordMap := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		wordMap[w] = true
	}
	if !wordMap[endWord] {
		return [][]string{}
	}
	curLayer := layerItem{
		result: [][]string{
			[]string{beginWord},
		},
		visitedMaps: []map[string]bool{
			map[string]bool{
				beginWord: true,
			},
		},
	}

	finded := false
	for len(curLayer.result) > 0 {
		tempLI := layerItem{
			result:      [][]string{},
			visitedMaps: []map[string]bool{},
		}
		rets := curLayer.result
		for index, ret := range rets {
			visitedMap := curLayer.visitedMaps[index]

			w := ret[len(ret)-1]
			cArr := []rune(w)
			for i, c := range cArr {
				for rc := 'a'; rc <= 'z'; rc++ {
					cArr[i] = rc
					tmpW := string(cArr)
					if !wordMap[tmpW] {
						cArr[i] = c
						continue
					}
					if visitedMap[tmpW] {
						cArr[i] = c
						continue
					}
					iret := append([]string{}, ret...)
					iret = append(iret, tmpW)
					ivm := copyMap(visitedMap)
					ivm[tmpW] = true

					tempLI.result = append(tempLI.result, iret)
					tempLI.visitedMaps = append(tempLI.visitedMaps, ivm)
					if tmpW == endWord {
						finded = true
					}
				}
			}
		}
		if finded {
			return tempLI.result
		}
		curLayer = tempLI
	}
	return [][]string{}
}

func copyMap(mp map[string]bool) map[string]bool {
	ret := make(map[string]bool, len(mp))
	for k, v := range mp {
		ret[k] = v
	}
	return ret
}

// @lc code=end

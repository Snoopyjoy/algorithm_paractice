package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start

var btnMap map[rune][]rune = map[rune][]rune{
	'2': []rune("abc"),
	'3': []rune("def"),
	'4': []rune("ghi"),
	'5': []rune("jkl"),
	'6': []rune("mno"),
	'7': []rune("pqrs"),
	'8': []rune("tuv"),
	'9': []rune("wxyz"),
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	runes := make([][]rune, 0, len(digits))
	for _, v := range digits {
		runes = append(runes, btnMap[v])
	}
	reto := combineRunes(runes)
	ret := make([]string, len(reto))
	for i, v := range reto {
		ret[i] = string(v)
	}
	return ret
}

func combineRunes(runes [][]rune) [][]rune {
	ret := [][]rune{}
	if len(runes) == 1 {
		for _, r := range runes[0] {
			ret = append(ret, []rune{r})
		}
		return ret
	}
	r := runes[0]
	child := combineRunes(runes[1:])
	for _, c := range r {
		for _, cc := range child {
			temp := make([]rune, 0, len(cc)+1)
			temp = append(temp, c)
			temp = append(temp, cc...)
			ret = append(ret, temp)
		}
	}
	return ret
}

// @lc code=end

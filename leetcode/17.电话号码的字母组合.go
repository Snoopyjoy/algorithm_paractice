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
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	btnMap := make(map[int]([]string), 8)
	btnMap[2] = []string{"a", "b", "c"}
	btnMap[3] = []string{"d", "e", "f"}
	btnMap[4] = []string{"g", "h", "i"}
	btnMap[5] = []string{"j", "k", "l"}
	btnMap[6] = []string{"m", "n", "o"}
	btnMap[7] = []string{"p", "q", "r", "s"}
	btnMap[8] = []string{"t", "u", "v"}
	btnMap[9] = []string{"w", "x", "y", "z"}

	digitsArr := strings.Split(digits, "")
	return combineLetter(btnMap, digitsArr)
}

func combineLetter(btnMap map[int]([]string), digitsArr []string) []string {
	result := []string{}
	if len(digitsArr) == 1 {
		d := digitsArr[0]
		num, _ := strconv.Atoi(d)
		letters := btnMap[num]
		return letters
	}
	d := digitsArr[0]
	leftDigits := digitsArr[1:]
	num, _ := strconv.Atoi(d)
	letters := btnMap[num]
	leftCombines := combineLetter(btnMap, leftDigits)
	for _, l := range letters {
		for _, c := range leftCombines {
			r := append([]string{l}, c)
			result = append(result, strings.Join(r, ""))
		}
	}
	return result
}

// @lc code=end

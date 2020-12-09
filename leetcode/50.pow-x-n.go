package leetcode_50

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if x == 0 || x == 1 {
		return x
	}
	if n == 0 {
		return 1
	}
	var f func(x float64) float64
	if n > 0 {
		f = fn
	} else {
		f = fn1
	}
	ret := float64(1)

	v := myPow(x, n/2)

	ret = v * v

	if n%2 != 0 {
		ret = ret * f(x)
	}
	return ret
}

func fn1(x float64) float64 {
	return 1 / x
}

func fn(x float64) float64 {
	return x
}

// @lc code=end

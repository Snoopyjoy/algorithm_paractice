/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	primes := make([]int, n)
	for i := 2; i*i < n; i++ {
		if primes[i] == 0 && isPrime(i) {
			for j := i * i; j < n; j += i {
				primes[j] = 1
			}
		}
	}
	ret := 0
	for i := 2; i < n; i++ {
		if primes[i] == 0 {
			ret++
		}
	}
	return ret
}

func isPrime(n int) bool {
	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// @lc code=end

var count int

func countPrimes1(n int) int {
	count = 0
	if n < 2 {
		return 0
	}
	origin, wait := make(chan int), make(chan struct{})
	process(origin, wait)
	for num := 2; num < n; num++ {
		origin <- num
	}
	close(origin)
	<-wait
	return count
}

func process(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		count++
		out := make(chan int)
		process(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}

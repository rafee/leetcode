/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

package golang

// @lc code=start
func countPrimes(n int) int {
	count := 0
	notPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
			for j := i * i; j < n; j += i {
				notPrime[j] = true
			}
		}
	}
	return count
}

// @lc code=end

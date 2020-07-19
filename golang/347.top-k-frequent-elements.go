/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

package golang

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	numSlice := make([][2]int, 0)
	for num, freq := range freqMap {
		numSlice = append(numSlice, [2]int{num, freq})
	}

	selectTopKNums(numSlice, 0, len(numSlice)-1, k)
	return []int{}
}

func selectTopKNums(numFreq [][2]int, low, high, k int) {
	pivot := low
	low++
	for i := low; i < high; i++ {
		if numFreq[low][1] > numFreq[pivot][1] {
			swapNums(pivot, i, numFreq)
			pivot++
		}
	}
	swapNums(pivot, high, numFreq)
	if pivot+1 == k {
		return
	} else if pivot+1 > k {
		selectTopKNums(numFreq, low, pivot-1, k-low)
	} else {
		selectTopKNums(numFreq, pivot+1, high, k)
	}
}

func swapNums(high int, low int, numFreq [][2]int) {
	numFreq[low], numFreq[high] = numFreq[high], numFreq[low]
}

// @lc code=end

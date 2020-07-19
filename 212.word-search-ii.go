/*
 * @lc app=leetcode id=212 lang=golang
 *
 * [212] Word Search II
 */

package golang

// @lc code=start

type wdObject struct {
	word string
	pos  int
}

func newWord(word string) wdObject {
	this := new(wdObject)
	this.word = word
	return *this
}

func (o wdObject) Len() int {
	return len(o.word)
}

func findWords(board [][]byte, words []string) []string {
	res := make([]string, 0)
	wdArr := make([]wdObject, len(words))
	for i, word := range words {
		wdArr[i] = newWord(word)
	}
	return res
}

// @lc code=end

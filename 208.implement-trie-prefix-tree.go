/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

package leetcode

// @lc code=start
type Trie struct {
	children map[byte]*Trie
	isWord   bool
}

/** Initialize your data structure here. */
func TrieConstructor() Trie {
	trie := new(Trie)
	trie.children = make(map[byte]*Trie)
	return *trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.isWord = true
		return
	}
	next, found := this.children[word[0]]
	if found {
		next.Insert(word[1:])
	} else {
		nextTrie := TrieConstructor()
		nextTrie.Insert(word[1:])
		this.children[word[0]] = &nextTrie
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.isWord
	}
	next, found := this.children[word[0]]
	if found {
		return next.Search(word[1:])
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	next, found := this.children[prefix[0]]
	if found {
		return next.StartsWith(prefix[1:])
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

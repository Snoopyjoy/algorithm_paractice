package leetcode

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	children map[rune]*Trie
	hasEnd   bool
}

/** Initialize your data structure here. */
func TrieConstructor() Trie {
	t := Trie{}
	t.children = map[rune]*Trie{}
	return t
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		if ct, ok := node.children[c]; ok {
			node = ct
		} else {
			newNode := TrieConstructor()
			node.children[c] = &newNode
			node = &newNode
		}
	}
	node.hasEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		if ct, ok := node.children[c]; ok {
			node = ct
		} else {
			return false
		}
	}
	return node.hasEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		if ct, ok := node.children[c]; ok {
			node = ct
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

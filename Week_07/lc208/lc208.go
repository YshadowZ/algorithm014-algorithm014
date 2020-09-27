package lc208

import (
	"strings"
)

type Trie struct {
	Value string
	Son   *map[string]*Trie
	End   bool // 是否是个单词
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Value: "*", Son: &map[string]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	Son := this.Son
	wordArr := strings.Split(word, "")
	index := 0
	var trie Trie
	for {
		if index == len(wordArr) {
			trie.End = true
			break
		}
		if (*Son)[wordArr[index]] != nil {
			trie = *((*Son)[wordArr[index]])
			Son = (*Son)[wordArr[index]].Son
			index++
		} else {
			trie = Trie{Value: wordArr[index], Son: &map[string]*Trie{}}
			(*Son)[wordArr[index]] = &trie
			Son = (*Son)[wordArr[index]].Son
			index++
		}
	}

	return
}

// Search Returns if the word is in the trie
func (this *Trie) Search(word string) bool {
	index := 0
	wordArr := strings.Split(word, "")
	length := len(wordArr)
	next := "*"
	Son := this.Son
	for {
		if (*Son)[wordArr[index]] != nil {
			next = (*Son)[wordArr[index]].Value
			Son = (*Son)[wordArr[index]].Son
			index++
		} else {
			next = ""
		}
		if next == "" {
			return false
		}
		if index == length && (*Son)[wordArr[index]].End {
			return (*Son)[wordArr[index]].End
		}
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.Search(prefix)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

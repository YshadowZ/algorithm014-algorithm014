package main

import (
	"week_07/lc208"
)

func main() {
	trie := lc208.Constructor()
	trie.Insert("apple")
	trie.Insert("appl")
	// result := trie.Search("appl")
	// fmt.Println(trie, result)
}

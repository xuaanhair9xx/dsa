package main

import "fmt"

const ALPHABET_SIZE = 26

type TrieNode struct {
	children []*TrieNode
	isEndOfWord bool
}

func getNode() *TrieNode {
	pNode := &TrieNode{}
	pNode.isEndOfWord = false
	pNode.children = make([]*TrieNode, ALPHABET_SIZE)
	return pNode
}

func (trie *TrieNode) Insert(key string) {
	for i := 0; i < len(key); i ++ {
		index := key[i] - 'a'
		if trie.children[index] == nil {
			trie.children[index] = getNode()
		}
		trie = trie.children[index]
	}
	trie.isEndOfWord = true
}

func (trie *TrieNode) Search(key string) bool {
	for _, v := range key {
		index := v - 'a'
		if trie.children[index] == nil {
			return false
		}
		trie = trie.children[index]
	}
	return trie.isEndOfWord;
}

func main() {
	keys := []string {"the", "a", "there", "answer", "any", "by", "bye", "their"};
	root := getNode()

	for _, key := range keys {
		root.Insert(key)
	}

	fmt.Println(root.Search("their"))
}
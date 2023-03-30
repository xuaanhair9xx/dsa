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
	board := [][]rune {
		{'o','a','a','n'},
		{'e','t','a','e'},
		{'i','h','k','r'},
		{'i','f','l','v'},
	}
	words := []string{"oath","pea","eat","rain"}
	way1(board, words)
}

func way1(board [][]rune, words []string) {
	var res []string
	var dfs func(int, int, string, *TrieNode)
	rows, cols := len(board), len(board[0])

	dfs = func(r, c int, word string, node *TrieNode) {
		if r < 0 || c < 0 || r >= rows || c >= cols {
			return
		}
		if node.children[board[r][c] - 'a'] == nil {
			return
		}
		word = word + string(board[r][c])
		node = node.children[board[r][c] - 'a']
		if node.isEndOfWord {
			res = append(res, word)
		}
		dfs(r+1, c, word, node)
		dfs(r-1, c, word, node)
		dfs(r, c+1, word, node)
		dfs(r, c-1, word, node)
	}

	root := getNode()
	for _, key := range words {
		root.Insert(key)
	}

	for i := 0; i < rows; i ++ {
		for j := 0; j < cols; j++ {
			dfs(i,j,"",root)
		}
	}

	fmt.Println(res)
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TrieNode struct {
	children map[string]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[string]*TrieNode),
		isEnd:    false,
	}
}

func main() {
	trie := NewTrieNode()
	file, _ := os.Open("input19.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	words := strings.Split(scanner.Text(), ", ")

	for _, word := range words {
		current := trie
		for _, char := range word {
			s := string(char)
			if _, ok := current.children[s]; !ok {
				current.children[s] = NewTrieNode()
			}
			current = current.children[s]
		}
		current.isEnd = true
	}

	scanner.Scan()

	count := 0
	for scanner.Scan() {
		s := scanner.Text()
		if canFormPattern(s, trie) {
			count++
		}
	}
	fmt.Println(count)
}

func canFormPattern(s string, trie *TrieNode) bool {
	if len(s) == 0 {
		return true
	}

	current := trie
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if next, ok := current.children[char]; ok {
			current = next
			if current.isEnd {
				if canFormPattern(s[i+1:], trie) {
					return true
				}
			}
		} else {
			break
		}
	}
	return current.isEnd && len(s) == 0
}

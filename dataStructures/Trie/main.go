package main

import (
	"fmt"
	"math"
)

const Size = 27

// Node represents each nodes in trie data structure
type Node struct {
	children [Size]*Node
	isEnd    bool
}

// Trie represents the root node
type Trie struct {
	root *Node
}

//initTrie creating the trie structure

func initTrie() *Trie {
	return &Trie{root: &Node{}}
}

//insert inserting a new word in trie structure

func (t *Trie) insert(word string) {
	currentNode := t.root
	for _, w := range word {
		charIndex := -1
		if w == 32 {
			charIndex = 26
		} else {
			charIndex = int(math.Abs(float64(w - 'a')))
		}
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// search searching for a word in trie
func (t *Trie) search(word string) bool {
	currentNode := t.root
	for _, w := range word {
		charIndex := -1
		if w == 32 {
			charIndex = 26
		} else {
			charIndex = int(math.Abs(float64(w - 'a')))
		}
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd {
		return true
	} else {
		return false
	}
}

func main() {
	t := initTrie()
	words := []string{
		"Hello", "Warm Welcome", "Hi", "Welcome", "Thank you", "you",
	}
	for _, word := range words {
		t.insert(word)
	}
	fmt.Println("word present :", t.search("Welcome"))
}

package main

import "log"

// AlphabetSize : is the number of possible characters in the trie
const AlphabetSize int = 26

// Node : represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd bool
}

// Trie : represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie : will create a new Trie
func InitTrie() *Trie {
	return &Trie{root:&Node{}}
}

// Insert : will take in a word and add it to the trie
func (t *Trie) Insert(word string) {
	currentNode := t.root

	for _, w := range word {
		charIndex := w - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}

		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

// Search : will take in a word and return true if that word is included in the trie
func (t *Trie) Search(word string) bool {
	currentNode := t.root

	for _, w := range word {
		charIndex := w - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.children[charIndex]
	}

	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main() {
	myTrie := InitTrie()
	log.Println(myTrie.root)

	toAdd := []string{"aragorn", "aragon", "argon", "eragon", "oregon", "oregano", "oreo", "ore"}
	
	for _, v := range toAdd{
		myTrie.Insert(v)
	}
	log.Println(myTrie.Search("oreo"))

}
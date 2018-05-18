// Package trie provides trie datastructure
package trie

const (
	// CharSize is maximum child node a single trie node can have.
	CharSize = 128
)

// Trie is the interface that wraps the basic trie operation
type Trie interface {
	// Insert key into trie node datastructure. TimeComplexity = O(k)
	Insert(k []byte)

	// search for key in datastructure. TimeComplexity = O(k)
	Search(k []byte) (found bool, final bool)
}

// Node represents TrieNode using fixed-size array. This representation might consume more space and is not recommended, please refer to TrieHashNode for hashImplementation.
type Node struct {
	// Child nodes
	Children [CharSize]*Node

	// True if Node is finalState in state-machine
	FinalState bool
}

// NewNode will return new trieNode
func NewNode() (tNode *Node) {
	tNode = &Node{FinalState: false}
	return
}

// Insert key into trie node datastructure
func (root *Node) Insert(key []byte) {
	var index byte

	tNode := root
	for level := 0; level < len(key); level++ {
		index = key[level]

		if tNode.Children[index] == nil {
			tNode.Children[index] = NewNode()
		}
		tNode = tNode.Children[index]
	}

	// mark last node as final node.
	tNode.FinalState = true
}

// Search for a key in trie node datastructure
func (root *Node) Search(key []byte) (found bool, final bool) {
	var index byte

	tNode := root
	for level := 0; level < len(key); level++ {
		index = key[level]

		if tNode.Children[index] == nil {
			return false, false
		}
		tNode = tNode.Children[index]
	}
	return tNode != nil, tNode.FinalState
}

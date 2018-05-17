// trie package provides trie datastructure
package trie

const (
	CHAR_SIZE = 128
)

type Trie interface {
	Insert(k []byte)                          // Insert key into trie node datastructure. TimeComplexity = O(k)
	Search(k []byte) (found bool, final bool) // search for key in datastructure. TimeComplexity = O(k)
}

// Trie representation using fixed-size array. This representation might consume more space in some cases and is not recommended, please refer to TrieHashNode for hashImplementation.
type TrieNode struct {
	Children [CHAR_SIZE]*TrieNode

	FinalState bool
}

func NewTrieNode() (tNode *TrieNode) {
	tNode = &TrieNode{FinalState: false}
	return
}

func (root *TrieNode) Insert(key []byte) {
	var index byte

	tNode := root
	for level := 0; level < len(key); level++ {
		index = key[level]

		if tNode.Children[index] == nil {
			tNode.Children[index] = NewTrieNode()
		}
		tNode = tNode.Children[index]
	}

	// mark last node as final node.
	tNode.FinalState = true
}

func (root *TrieNode) Search(key []byte) (found bool, final bool) {
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

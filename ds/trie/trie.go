// trie package provides trie datastructure
package trie

const (
	CHAR_SIZE = 128
)

type Trie interface {
	Insert(k []byte)                          // Insert key into trinode datastructure. Complexity O(k)
	Search(k []byte) (found bool, final bool) // search for key in datastructure. Complexity O(k). found=ngram-substringMatch?, final=exactMatch?
}

// Trie ds implemented with fixed size array. This may consume more space. Refer to TrieHashNode that implements hashmap (hashtable) instead of array.
type TrieNode struct {
	Children [CHAR_SIZE]*TrieNode

	FinalState bool
}

// NewNode returns new TrieNode
func NewTrieNode() (tNode *TrieNode) {
	tNode = &TrieNode{FinalState: false}
	return
}

// Insert key into trieNode if not present
// If the key is prefix of trie node, just marks final/leaf node.
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

// Search for key in trie
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

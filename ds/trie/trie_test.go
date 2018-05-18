package trie

import "testing"
import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestNode(t *testing.T) {
	trieNode := NewNode()

	if trieNode.FinalState {
		t.Error("New Node should not be final.")
	}

	trieNode.Insert([]byte{122, 23, 55, 124})

	trieNode.Insert([]byte{14, 33, 42})

	if trieNode.Children[14] == nil || trieNode.Children[122] == nil {
		t.Error("Insertion Failed")
	}

	r, final := trieNode.Search([]byte{14, 33, 42})
	if !(r && final) {
		t.Error("Should found the key 1")
	}

	r, final = trieNode.Search([]byte{14, 33})
	if !r {
		t.Error("Trie has all nodes from input string")
	}
	if final {
		t.Error("Key doesnot exist but found in search")
	}

	r, final = trieNode.Search([]byte{14, 3})
	if r {
		t.Error("Trie has all nodes from input string")
	}
	if final {
		t.Error("Key doesnot exist but found in search")
	}
}

func TestHashNode(t *testing.T) {
	trieNode := NewHashNode()

	if trieNode.FinalState {
		t.Error("New Node should not be final.")
	}

	trieNode.Insert([]byte{122, 23, 55, 124})

	trieNode.Insert([]byte{14, 33, 42})

	if trieNode.Children[14] == nil || trieNode.Children[122] == nil {
		t.Error("Insertion Failed")
	}

	r, final := trieNode.Search([]byte{14, 33, 42})
	if !(r && final) {
		t.Error("Should found the key 1")
	}

	// for given trie, leaf(33) exist in tree so the input keyword is found but it doesnot reached final state.
	r, final = trieNode.Search([]byte{14, 33})
	if !r {
		t.Error("Trie has all nodes from input string")
	}
	if final {
		t.Error("Key doesnot exist but found in search")
	}

	// for given trie, leaf(3) doesnot exist in tree so the input keyword is not found nor it reached final state.
	r, final = trieNode.Search([]byte{14, 3})
	if r {
		t.Error("Trie has all nodes from input string")
	}
	if final {
		t.Error("Key doesnot exist but found in search")
	}

}

func BenchmarkNodeInsert(b *testing.B) {
	tNode := NewNode()
	for i := 0; i < b.N; i++ {
		a := makeByteArray(5)
		tNode.Insert(a)
	}
}

func BenchmarkNodeSearch(b *testing.B) {
	tNode := NewNode()
	for i := 0; i < b.N; i++ {
		c := makeByteArray(5)
		tNode.Search(c)
	}
}

func BenchmarkHashNodeInsert(b *testing.B) {
	thNode := NewHashNode()
	for i := 0; i < b.N; i++ {
		a := makeByteArray(5)
		thNode.Insert(a)
	}
}

func BenchmarkHashNodeSearch(b *testing.B) {
	tNode := NewHashNode()
	for i := 0; i < b.N; i++ {
		c := makeByteArray(5)
		tNode.Search(c)
	}
}

func makeByteArray(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

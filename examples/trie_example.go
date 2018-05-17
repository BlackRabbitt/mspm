package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"time"

	. "github.com/BlackRabbitt/mspm/ds/trie"
)

var input = `
google chrome
microsoft
ibm
google glass
apple
ball
`

func main() {
	tNode := NewTrieHashNode()

	scanner := bufio.NewScanner(strings.NewReader(input))

	start := time.Now()

	for scanner.Scan() {
		tNode.Insert(scanner.Bytes())
	}
	log.Println("took ", time.Since(start), " to insert input")

	t := []byte("google chrome")
	res := tNode.Search(t)
	log.Println(res)
}

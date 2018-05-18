# Multi-String Pattern Matching algorithm.
[![Go Report Card](https://goreportcard.com/badge/github.com/BlackRabbitt/mspm)](https://goreportcard.com/report/github.com/BlackRabbitt/mspm)

This implementation is inspired from [Aho-Corasick algorithm](https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm)

[GoDoc Api Documentation](https://godoc.org/github.com/BlackRabbitt/mspm)

## Getting Started
```
modelA = mspm.NewModel("mspm_model_A")
patternsToSearch = strings.NewReader(words) // words is newline seperated list of words/keywords

// build mspm model with patterns
modelA.Build(patternsToSearch)

inputString := "input document where the above pattern is searched for."
document := strings.NewReader(inputString)
output, err := modelA.MultiTermMatch(document)

// output ~= [{matched_word: n_count}, ..]
```

## Test Coverage
* [trie package](https://gocover.io/github.com/blackrabbitt/mspm/ds/trie)
* [mspm package](https://gocover.io/github.com/blackrabbitt/mspm/search)

## TrieNode vs TrieHashNode benchmark
```
$ cd github.com/BlackRabbitt/mspm/ds/trie
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/BlackRabbitt/mspm/ds/trie
BenchmarkTrieNodeInsert-4             500000          2582 ns/op
BenchmarkTrieNodeSearch-4           10000000           205 ns/op
BenchmarkTrieHashNodeInsert-4        1000000          1491 ns/op
BenchmarkTrieHashNodeSearch-4       10000000           206 ns/op
PASS
ok      github.com/BlackRabbitt/mspm/ds/trie	8.795s
```

## Resources
1. [Trie](https://en.wikipedia.org/wiki/Trie)
2. [mspm](http://www.ijsrp.org/research_paper_jul2012/ijsrp-july-2012-101.pdf) - Multi-String Pattern Matching algorithm. Generally used for Information Retrieval.
3. [Aho-Corasick algorithm](https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm)

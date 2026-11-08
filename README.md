# mspm

`mspm` is a Go library for matching a set of literal byte-string patterns in a document and counting the matches. It uses a trie and is inspired by the [Aho–Corasick algorithm](https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm).

> **Status:** This project is being revived. The current implementation is small and has known limitations; review the behavior below before using it for production workloads.

## Requirements

- Go 1.27 or newer

## Install

```sh
go get github.com/BlackRabbitt/mspm
```

The matcher API currently lives in the `search` package:

```go
import "github.com/BlackRabbitt/mspm/search"
```

## Example

```go
package main

import (
	"fmt"
	"strings"

	"github.com/BlackRabbitt/mspm/search"
)

func main() {
	model := search.NewModel("keywords")
	model.Build(strings.NewReader("mspm\npattern\nsearch\n"))

	counts, err := model.MultiTermMatch(strings.NewReader("mspm pattern search mspm"))
	if err != nil {
		panic(err)
	}

	fmt.Println(counts["mspm"])    // 2
	fmt.Println(counts["pattern"]) // 1
	fmt.Println(counts["search"])  // 1
}
```

Patterns are supplied one per line to `Build`. Matching is case-sensitive. `MultiTermMatch` reads the document and returns counts keyed by the matched pattern.

## Current behavior and limitations

- The implementation works on bytes, so it does not provide Unicode-aware character matching.
- Matches are consumed as the scan proceeds; overlapping matches are not all reported.
- The matcher is trie-based, but does not currently implement the failure links used by the full Aho–Corasick algorithm.
- `Build` does not return scanner errors, so unusually long pattern lines may not be reported to the caller.

These constraints are part of the current implementation and may change as the project is improved.

## Development

Run the package tests with:

```sh
go test ./...
```

## License

BSD-3-Clause. See [LICENSE](LICENSE).

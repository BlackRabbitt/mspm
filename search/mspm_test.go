package mspm

import (
	"strings"
	"testing"
)

var patterns = `
mspm
2
another
example
case
ball`

func TestMultiTermMatch(t *testing.T) {
	modelA := NewModel("Test1")
	words := strings.NewReader(patterns)
	modelA.Build(words)

	inputString := "mspm-algorithm is implemented to solve the problem of searching multiple keywords in an paragraph. For example in this test, if we search for mspm in this input string, the algorithm should return 2. Another TBD important thing is this algorithm is case-sensitive, i.e., if I search for 'Case' in this document then I will get 0 as an output but if I search for 'case' it will return me 2 instead of 4(ignoring Case)"
	document := strings.NewReader(inputString)
	output, err := modelA.MultiTermMatch(document)

	if err != nil {
		t.Log("Error must be nil")
	}

	if 4 != len(output) {
		t.Log("Test length of output")
		t.Errorf("Expected: %d, Got: %d", 4, len(output))
		t.Log(output)
	}

	if 2 != output["mspm"] {
		t.Log("Test count of 'mspm' keyword appearing in document")
		t.Errorf("Expected: %d, Got: %d", 2, output["mspm"])
	}

	if 2 != output["2"] {
		t.Log("Test count of number '2' keyword appearing in document")
		t.Errorf("Expected: %d, Got: %d", 2, output["2"])
	}

	if 1 != output["example"] {
		t.Log("Test count of 'example' keyword appearing in document")
		t.Errorf("Expected: %d, Got: %d", 2, output["example"])
	}

	if 2 != output["case"] {
		t.Log("Test for case sensitive")
		t.Errorf("Expected: %d, Got: %d", 2, output["case"])
	}

	if 0 != output["Case"] {
		t.Log("Test for case sensitive")
		t.Errorf("Expected: %d, Got: %d", 0, output["Case"])
	}
}

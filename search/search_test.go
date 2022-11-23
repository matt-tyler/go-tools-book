package search_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matt-tyler/go-tools-book/search"
)

func TestLines(t *testing.T) {
	t.Parallel()

	text := `My test input
is here split across many lines
o what lines shall I find -
will there ever be a rainline?
what dreams may come?
`

	query := "line"

	want := `is here split across many lines
o what lines shall I find -
will there ever be a rainline?
`

	input := bytes.NewBufferString(text)
	output := new(bytes.Buffer)

	searcher, err := search.New(
		search.WithInput(input),
		search.WithOutput(output))
	if err != nil {
		t.Fatal(err)
	}

	searcher.Lines(query)

	got := output.String()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Error(diff)
	}
}

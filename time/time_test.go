package time_test

import (
	"bytes"
	"testing"

	"time"

	timePkg "github.com/matt-tyler/go-tools-book/time/i18n"
)

func TestPrintsTimeToWriter(t *testing.T) {
	t.Parallel()

	fakeTerminal := &bytes.Buffer{}

	currentTime := time.Date(
		2009, 11, 17, 20, 1, 58, 651387237, time.UTC)

	// timePkg.PrintTo(fakeTerminal, currentTime)
	timePkg.PrintTo(fakeTerminal, "en-gb", currentTime)

	want := "It's 1 minute past 8"
	got := fakeTerminal.String()

	if got != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}

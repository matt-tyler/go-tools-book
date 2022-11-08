package simple

import (
	"fmt"
	"io"
	"time"
)

func layout(t time.Time) string {
	if t.Minute() == 1 {
		return "It's 4 minute past 3"
	}
	return "It's 4 minutes past 3"
}

func PrintTo(w io.Writer, t time.Time) {
	layout := layout(t)
	fmt.Fprint(w, t.Format(layout))
}

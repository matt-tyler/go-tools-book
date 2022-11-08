package plural

import (
	"io"
	"sync"
	"time"

	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var onlyOnce sync.Once
var p = message.NewPrinter(language.English)

func PrintTo(w io.Writer, t time.Time) {
	onlyOnce.Do(func() {
		message.Set(language.English, "It's %d minute(s) past %d",
			plural.Selectf(1, "%d",
				plural.One, "It's %[1]d minute past %[2]d",
				plural.Other, "It's %[1]d minutes past %[2]d",
			))
	})
	p.Fprintf(w, "It's %d minute(s) past %d", t.Minute(), t.Hour()%12)
}

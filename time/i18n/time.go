package i18n

import (
	"errors"
	"io"
	"time"

	_ "github.com/matt-tyler/go-tools-book/internal/translations"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func getLanguage(locale string) (lang language.Tag, err error) {
	switch locale {
	case "en-gb":
		lang = language.MustParse("en-GB")
	case "de-de":
		lang = language.MustParse("de-DE")
	default:
		err = errors.New("unknown locale supplied")
		return lang, err
	}
	return lang, nil
}

func PrintTo(w io.Writer, locale string, t time.Time) error {
	lang, err := getLanguage(locale)
	if err != nil {
		return err
	}

	p := message.NewPrinter(lang)
	p.Fprintf(w, "It's %d minute(s) past %d", t.Minute(), t.Hour()%12)
	return nil
}

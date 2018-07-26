package word

import (
	"saltlab/foundation/rest"
	"saltlab/nba/def"
	"unicode"
	"unicode/utf8"
)

// CheckNick : check nick name.
func CheckNick(name string) error {
	len := utf8.RuneCountInString(name)
	if len < 3 || len > 10 {
		return rest.NewClientError(def.ErrInvalidNickSize)
	}

	for _, r := range name {
		if unicode.IsDigit(r) ||
			unicode.Is(unicode.Latin, r) ||
			unicode.Is(unicode.Hangul, r) ||
			r == '-' ||
			r == '_' {
			continue
		}
		return rest.NewClientError(def.ErrInvalidNickChar)
	}
	if IsInclude(name) {
		return rest.NewClientError(def.ErrInvalidNickWord)
	}
	return nil
}

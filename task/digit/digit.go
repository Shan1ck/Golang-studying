package digit

import (
	"strings"

	"main3.1/wordz"
)

func Digit() string {
	return strings.TrimPrefix(wordz.Random(), "Random word: ")
}

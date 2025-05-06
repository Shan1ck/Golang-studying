package city

import (
	"strings"

	"main3.1/wordz"
)

func City() string {
	cities := map[string]string{
		"One":   "Moscow",
		"Two":   "Chernobyl",
		"Three": "Novosibirsk",
		"Four":  "Krasnoyarsk",
		"Five":  "Peter",
	}
	id := strings.TrimPrefix(wordz.Random(), "Random word: ")
	return cities[id]
}

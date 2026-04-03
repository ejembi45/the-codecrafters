// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Faith Ejembi]
// Squad:  [Goroutines]

package modulestringtransformer

import (
	"strings"
	"unicode"
)

func cleantext(word string) string {
	var result strings.Builder

	for _, char := range word {
		//  this says only keep only letters, digits, and spaces
		if unicode.IsLetter(char) || unicode.IsDigit(char) || char == ' ' {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func snakeCase(word string) string {
	word = strings.ToLower(word)
	word = cleantext(word)

	words := strings.Fields(word)

	// removes  the "snake" if it's the first word
	if len(words) > 0 && words[0] == "snake" {
		words = words[1:]
	}

	return strings.Join(words, "_")
}

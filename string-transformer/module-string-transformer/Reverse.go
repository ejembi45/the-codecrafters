// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Faith Ejembi]
// Squad:  [Goroutines]

package modulestringtransformer

import (
	"strings"
)

func Reverse(word string) string {
	word = strings.ToLower(word)

	if strings.HasPrefix(word, "reverse") && strings.ToLower(word) == "reverse" {
		strings.TrimPrefix(word, "reverse")
	}

	runes := []rune(word)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func ReversedText(input string) string {
	words := strings.Fields(input)

	reversedWords := []string{}

	for _, word := range words {
		reversedWords = append(reversedWords, Reverse(word))
	}

	return strings.Join(reversedWords, " ")
}

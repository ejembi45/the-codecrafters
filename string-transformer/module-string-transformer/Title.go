// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Faith Ejembi]
// Squad:  [Goroutines]

package modulestringtransformer

import (
	"strings"
)

func Title(word string) string {
	word = strings.ToLower(word)
	words := strings.Fields(word)
	smallwords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true, "but": true,
		"or": true, "for": true, "nor": true, "on": true, "at": true,
		"to": true, "by": true, "in": true, "of": true, "up": true,
		"as": true, "is": true, "it": true,
	}

	//  This Removes leading "title" if present
	if len(words) > 0 && words[0] == "title" {
		words = words[1:]
	}

	for i, w := range words {
		if i == 0 || !smallwords[w] {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}

	return strings.Join(words, " ")
}

// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Faith Ejembi]
// Squad:  [Goroutines]

package modulestringtransformer

import (
	"strings"
)

func ToLower(word string) string {
	if strings.HasPrefix(word, "lower") {
		if len(word) > 0 {
			word = strings.TrimPrefix(word, "lower")
			word = strings.ToLower(word)
		}
	}
	return word
}

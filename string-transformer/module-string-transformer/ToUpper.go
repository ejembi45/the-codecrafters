// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Faith Ejembi]
// Squad:  [Goroutines]

package modulestringtransformer

import (
	"strings"
)

func ToUpper(word string) string {
	if strings.HasPrefix(word, "upper") {
		word = strings.TrimPrefix(word, "upper")
		word = strings.ToUpper(word)
	}
	return word
}

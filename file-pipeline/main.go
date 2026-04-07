// ═══════════════════════════════════════════
// SQUAD: PIPELINE CONTRACT
// Squad: Goroutines
// Name: Faith Ochanya Ejembi
// ───────────────────────────────────────────
//
// Input line types:
//
//   1. Normal report lines
//   2. Lines in ALL CAPS
//   3. Lines in all lowercase
//   4. Lines starting with TODO:
//   5. Lines starting with CLASSIFIED:
//   6. Lines that are only dashes or blank
//   7. Lines with leading/trailing whitespace
//   8. Lines containing numbers and symbols
//
// Transformation rules (in order):
//
//   1. Trim all leading and trailing whitespace
//   2. Remove lines that are only dashes or blank
//   3. Replace TODO: with ✦ ACTION:
//   4. Replace CLASSIFIED: with [REDACTED]:
//   5. Reverse the words in any line that contains the word REVERSE
//
// Output format:
//
//   Header: yes — SENTINEL FIELD REPORT — PROCESSED
//   Line numbering format: 001., 002., 003. (three-digit zero-padded)
//   Summary block: yes — Lines Processed, Lines Written, Lines Removed
//
// Terminal summary fields:
//
//   ✦ Lines read    : <number>
//   ✦ Lines written : <number>
//   ✦ Lines removed : <number>
//   ✦ Rules applied : Trim whitespace, Remove blank/dash lines, Replace TODO, Replace CLASSIFIED, Reverse REVERSE lines
// ═══════════════════════════════════════════

package main

import (
	"strings"
)

func reverseWords(s string) string {
	if strings.Contains(s, "REVERSE") {
		words := strings.Fields(s)
		for i, n := 0, len(words)-1; i < n; i, n = i+1, n-1 {
			words[i], words[n] = words[n], words[i]
		}
		return strings.Join(words, " ")
	}
	return s
}

func lowerToUpper(word string) string {
	if word == strings.ToLower(word) {
		return strings.ToUpper(word)
	}
	return word
}

func trimWhitespace(word string) string {
	return strings.TrimSpace(word)
}

func replaceTodo(word string) string {
	return strings.ReplaceAll(word, "TODO:", " ACTION:")
}

func replaceClassified(word string) string {
	if strings.HasPrefix(word, "CLASSIFIED:") {
		return "[REDACTED]:" + word[len("CLASSIFIED:"):]
	}
	return word
}

func main() {

}

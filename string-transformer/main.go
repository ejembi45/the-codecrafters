// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Faith Ejembi]
// Squad:  [Goroutines]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func ToCapital(word string) string {
	s := strings.ToLower(word) //converts the text to lowercase
	strings.Fields(word)
	// As the name implies prefix: it checks for the prefix before execution
	if strings.HasPrefix(s, "cap") {
		s = strings.TrimPrefix(s, "cap")
		if len(s) > 0 {
			s = strings.Title(s)
		}
	}
	return s
}

func ToLower(word string) string {
	word = strings.ToLower(word)
	if strings.HasPrefix(word, "lower") {
		if len(word) > 0 {
			word = strings.TrimPrefix(word, "lower") // removes lower from the output
			word = strings.ToLower(word)             //converts any text to lower case
		}
	}
	return word
}

func ToUpper(word string) string {
	word = strings.ToLower(word)
	if strings.HasPrefix(word, "upper") {
		word = strings.TrimPrefix(word, "upper")
		word = strings.ToUpper(word)
	}
	return word
}

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

func snakeCase(text string) string {
	words := strings.Fields(strings.ToLower(text))

	if len(words) > 0 && words[0] == "snake" {
		words = words[1:]
	}

	return strings.Join(words, "_")
}

func Reverse(word string) string {
	runes := []rune(word)
	for i, n := 0, len(runes)-1; i < n; i, n = i+1, n-1 {
		runes[i], runes[n] = runes[n], runes[i]
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

func showHelp() {
	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║     SENTINEL STRING TRANSFORMER              ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Println()

	fmt.Println("USAGE:")
	fmt.Println("  <command> <text>")
	fmt.Println()

	fmt.Println("COMMANDS:")
	fmt.Println("  cap       → Capitalize each word")
	fmt.Println("  lower     → Convert text to lowercase")
	fmt.Println("  upper     → Convert text to uppercase")
	fmt.Println("  title     → Smart title formatting")
	fmt.Println("  snake     → Convert text to snake_case")
	fmt.Println("  reverse   → Reverse each word")
	fmt.Println("  exit      → Quit the program")
	fmt.Println()

	fmt.Println("EXAMPLES:")
	fmt.Println("  > cap director adaeze okonkwo")
	fmt.Println("    Director Adaeze Okonkwo")
	fmt.Println()

	fmt.Println("  > snake Hello World Program")
	fmt.Println("    hello_world_program")
	fmt.Println()

	fmt.Println("  > exit")
	fmt.Println("    Shutting down String Transformer...")
	fmt.Println()

	fmt.Println("──────────────────────────────────────────────")
	fmt.Println("Tip: Commands are case-sensitive, wrong input leads to error ")
}

func main() {
	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")
	fmt.Println("──────────────────────────────────────")
	fmt.Println("Input 'Help' ")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)

		word := strings.Fields(input)
		if len(word) == 0 {
			continue
		}

		command := strings.ToLower(word[0])

		if len(word) == 1 && command != "help" && command != "exit" {
			fmt.Printf("No text provided. usage: %s <text>\n", command)
			continue
		}

		switch command {
		case "help":
			showHelp()
		case "upper", "Upper", "UPPER":
			fmt.Println(ToUpper(input))
		case "lower", "Lower", "LOWER":
			fmt.Println(ToLower(input))
		case "cap", "Cap", "CAP":
			fmt.Println(ToCapital(input))
		case "title", "Title", "TITLE":
			fmt.Println(Title(input))
		case "snake", "Snake", "SNAKE":
			fmt.Println(snakeCase(input))
		case "reverse", "Reverse", "REVERSE":
			fmt.Println(ReversedText(input))
		case "exit", "Exit", "EXIT":
			fmt.Println("Shutting down String Transformer. Good bye")
			return
		default:
			fmt.Printf("Invalide command:\"%s\"\n", command)
			fmt.Println("valid commands: upper,lower,cap,title,snake,reverse,exit")
		}
	}

}

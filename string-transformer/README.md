
# String Transformer — Operation Gopher Protocol

## What It Does

The String Transformer is a command-line tool that takes corrupted text and
cleans it up using one of six transformation commands. You type a command
followed by your text, and the program outputs the corrected result. It keeps
running until you type "exit".

## How to Run

Clone the repo and move to the project folder:

    git clone the repository on github
    cd the-codecrafters/string-transformer

Then  u will run 

    go run main.go

## Commands and Examples

**upper** — converts all letters to uppercase

    > upper sentinel is watching
      → SENTINEL IS WATCHING

**lower** — converts all letters to lowercase

    > lower ALERT LEVEL FIVE DETECTED
      → alert level five detected

**cap** — capitalises the first letter of every word, rest go lowercase

    > cap THREAT LEVEL elevated
      → Threat Level Elevated

**title** — like cap, but small connector words stay lowercase unless
they are the first word. Small words: a, an, the, and, but, or, for,
nor, on, at, to, by, in, of, up, as, is, it

    > title the fall of the western power grid
      → The Fall of the Western Power Grid

**snake** — converts text to snake_case. All lowercase, spaces become
underscores. Any character that is not a letter, digit, or underscore
is removed.

    > snake Alert! Level 5 detected.
      → alert_level_5_detected

**reverse** — reverses each word individually. Word order stays the same.

    > reverse Go saves the world
      → oG sevas eht dlrow

## Edge Cases  to Handle

1.Unknown commands print an error and list valid options
2.Missing text after a command prints a usage message
3.Leading, trailing, and extra internal spaces are cleaned before processing
4.Numbers and symbols pass through without crashing
5.Empty input lines are ignored  the prompt just reappears
6.Commands are case-insensitive "UPPER", "Upper", and "upper" all work
7.For "title" the first word is always capitalised even if it is a small word
8.For "snake", only letters, digits, and underscores are kept everything
  else (punctuation, symbols) is removed before converting spaces to "_"

## THINK ABOUT

 1. Each transformation in my program is implemented as its own function e.g. upper, lower, cap, title, snake, reverse, which keeps the code modular and easy to maintain. The main() function is only responsible for handling user input and routing commands.
2.For the cap and title transformations, I split the input into words using strings.Fields. For each word, I isolate the first letter using slicing (word[:1]) and convert it to uppercase, while keeping the rest of the word unchanged or normalized as needed. The words are then joined back together using strings.Join.
3.In the title function, I store small words (such as "a", "the", "and") in a map[string]bool for efficient lookup. While iterating through the words, I check if the current word is the first word (i == 0). If it is, I always capitalize it. Otherwise, I only capitalize words that are not in the small words map. This ensures the title formatting rules are applied cleanly and efficiently.
# Base Converter — CodeCrafters CLI Tool

## What It Does

A command-line tool that converts numbers between bases.
Type a number and its base, and the program outputs the result.
It keeps running until you type "quit".

"hex" input to decimal output
 "bin" input to decimal output
 "dec" input to binary and hex output

## How to Run

    git clone the repository
    cd the-codecrafters/thecodecrafterthon-day2

    go run main.go

## Examples

**Hex to decimal:**

    > convert 1E hex
      ✦ Decimal: 30

**Binary to decimal:**

    > convert 10 bin
      ✦ Decimal: 2

**Decimal to binary and hex:**

    > convert 255 dec
      ✦ Binary: 11111111
      ✦ Hex:    FF

**Negative decimal:**

    > convert -12 dec
      ✦ Binary: -1100
      ✦ Hex:    -C

## Validation

The program handles bad input without crashing:

 "1G hex" → invalid hex digit
 "10201 bin" → invalid binary digit
"abc dec" → not a valid decimal number
 Missing arguments → usage message
 Empty input → prompt reappears silently

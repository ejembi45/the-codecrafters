//

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Base Converter (type 'quit' to exit)")

	for {
		var command, value, base string

		fmt.Print("> ")
		fmt.Scanln(&command, &value, &base)

		if command == "quit" {
			break
		}

		if command != "convert" {
			fmt.Println(" Unknown command")
			continue
		}

		value = strings.TrimSpace(value)
		base = strings.ToLower(strings.TrimSpace(base))

		if value == "" || base == "" {
			fmt.Println(" Invalid input")
			continue
		}

		switch base {
		case "hex":
			num, err := strconv.ParseInt(value, 16, 64)
			if err != nil {
				fmt.Println(" Invalid hex number")
				continue
			}
			fmt.Println(" Decimal:", num)

		case "bin":
			num, err := strconv.ParseInt(value, 2, 64)
			if err != nil {
				fmt.Println(" Invalid binary number")
				continue
			}
			fmt.Println(" Decimal:", num)

		case "dec":
			num, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				fmt.Println(" Invalid decimal number")
				continue
			}
			fmt.Println(" Binary: ", strconv.FormatInt(num, 2))
			fmt.Println(" Hex:    ", strings.ToUpper(strconv.FormatInt(num, 16)))

		default:
			fmt.Println(" Unknown base ( please input hexadecimal, binary, or decimal)")
		}
	}
}

package main

import (
	"ascii-art/ascii"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get all command-line arguments except the program name.
	// os.Args[0] is the program name, so it start from index 1.
	argSlice := os.Args[1:]

	// Validate if only one argument was passed, if not print an error message.
	if len(argSlice) < 1 || len(argSlice) > 1 {
		fmt.Println("This program requires only one argument! Your command contains less or more than one argument.")
		return
	}

	// Extract from argument slice (argSlice), the argument passed to the program.
	inputStr := argSlice[0]

	// If input string is empty, terminate the program
	if inputStr == "" {
		return
	}

	// Handle edge case where user passes "\n" as argument
	if inputStr == "\\n" {
		fmt.Println()
		return
	}

	// Split input string into multiple individual string using "\n" as the separator.
	// This enables the program to print a new line, if it is part of the argument.
	arrayOfStrings := strings.Split(inputStr, "\\n")

	// Read the ASCII art file. It contains the ascii representations for characters.
	asciiFile := ascii.ReadFile("standard.txt")

	// Print ascii representation of each string input
	ascii.AsciiPrinter(arrayOfStrings, asciiFile)
}

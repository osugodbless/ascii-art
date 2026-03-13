package ascii

import (
	"os"
	"strings"
)

// ReadFile reads a file from the given path and returns its contents
// as a slice of strings, where each element represents a line in the file.
func ReadFile(s string) ([]string, error) {
	// Read the entire file into memory
	file, err := os.ReadFile(s)

	// If an error occurs while reading the file, return the error
	if err != nil {
		return nil, err
	}

	// Split the file content by newline characters into a slice of strings
	asciiFile := strings.Split(string(file), "\n")

	// Return the slice of lines
	return asciiFile, nil
}

// AsciiPrinter converts input text into ASCII art using a banner file.
// arr: slice of input strings to print
// fileInput: slice containing ASCII art representations of characters
func AsciiPrinter(arr []string, fileInput []string) string {
	result := ""

	// Loop through each string in the input array
	for _, str := range arr {

		// If the string is empty, add a newline to preserve spacing
		if str == "" {
			result += "\n"
			continue
		}

		// Each ASCII character is represented by 8 lines
		for i := 1; i <= 8; i++ {

			// Loop through each character in the string
			for _, ch := range str {

				// Ensure the character is a printable ASCII character
				if ch >= 32 && ch <= 126 {

					// Calculate the starting index of the ASCII art
					// Each character occupies 9 lines in the banner file
					start := int((ch-32)*9) + i

					// Append the corresponding ASCII art line
					result += fileInput[start]
				}

			}

			// Add a newline after each ASCII art row
			result += "\n"
		}

	}

	// Return the final ASCII art result
	return result
}

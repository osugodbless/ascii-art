package ascii

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(s string) []string {
	file, err := os.ReadFile(s)

	if err != nil {
		fmt.Println("Unable to read file. Please check that your file path is correct and your directory contains the the file")
		return []string{}
	}

	asciiFile := strings.Split(string(file), "\n")

	return asciiFile
}

func AsciiPrinter(arr []string, input []string) {

	for _, str := range arr {
		if str == "" {
			fmt.Println()
			continue
		}

		for i := 1; i <= 8; i++ {
			for _, ch := range str {
				if ch >= 32 && ch <= 126 {
					start := int((ch-32)*9) + i
					fmt.Print(input[start])
				}

			}
			fmt.Println()

		}

	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	argSlice := os.Args[1:]

	if len(argSlice) < 1 || len(argSlice) > 1 {
		fmt.Println("This program requires only one argument! Your command contains less or more than one argument.")
		return
	}
	arg := argSlice[0]

	arg = strings.ReplaceAll(arg, "\\n", "\n")

	arrArg := strings.Split(arg, "\n")

	fmt.Println(arrArg)

	file, err := os.ReadFile("standard.txt")

	if err != nil {
		fmt.Println("Unable to read file. Please check that your file path is correct and your directory contains the the file")
		return
	}

	fileStr := strings.Split(string(file), "\n")

	// loop that handles printing of each row on a new line
	for _, str := range arrArg {
		if str == "" {
			fmt.Println()
		}

		for i := 1; i <= 8; i++ {
			for _, ch := range str {
				if ch >= 32 && ch <= 126 {
					start := int((ch-32)*9) + i
					fmt.Print(fileStr[start])
					// for j := start; j < start+1; j++ {
					// 	fmt.Print(fileStr[j])

					// }
				}

			}
			fmt.Println()

		}

	}

}

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(" >")
		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {

			break
		} else {
			cleanInput(text)
		}
		fmt.Println("echoing: ", text)
	}
}


func cleanInput(str string) []string {

	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	fmt.Println(words)

	return words
}
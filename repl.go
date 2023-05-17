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
			cleaned := cleanInput(text)


			if len(cleaned) == 0 {

				continue
			}

			command := cleaned[0]

			switch command {

			case "exit":
				os.Exit(0)

			case "help":
				fmt.Println("Welcome to the Pokedex help menu!")
				fmt.Println("Here are the available commands:")
				fmt.Println("help - get help")
				fmt.Println("exit - exit the program")
				fmt.Println("")

			default:
				fmt.Println("Invalid command")
				
			}
			
	
		}
}



type cliCommand struct {

	name string
	description string
	
}


func cleanInput(str string) []string {

	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
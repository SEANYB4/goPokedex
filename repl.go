package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
			fmt.Println(" >")
			scanner.Scan()
			text := scanner.Text()
			cleaned := cleanInput(text)


			if len(cleaned) == 0 {

				continue
			}

			commandName := cleaned[0]

			availableCommands := getCommands()

			command, ok := availableCommands[commandName]

			if !ok {
				fmt.Println("Invalid command")
				continue
			}
			
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
	
		}
}



type cliCommand struct {

	name string
	description string
	callback func(*config) error

}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand {

		"help" : {
			name: "help",
			description: "Prints the help menu",
			callback: callbackHelp,
		},
		"exit" : {
			name: "exit",
			description: "Turns off the Pokedex",
			callback: callbackExit,
		},

		"map" : {
			name: "map",
			description: "Displays available locations to travel to",
			callback: callbackMap,
		},
		"mapb" : {
			name: "mapb",
			description: "Displays previous available locations to travel to",
			callback: callbackMapB,
		},


	}
}


func cleanInput(str string) []string {

	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
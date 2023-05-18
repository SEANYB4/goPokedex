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
			fmt.Println()
			fmt.Print(" >")
			scanner.Scan()
			text := scanner.Text()
			cleaned := cleanInput(text)
			fmt.Println()


			if len(cleaned) == 0 {

				continue
			}

			commandName := cleaned[0]

			args := []string{}
			if len(cleaned) > 1 {
				args = cleaned[1:]
			}

			availableCommands := getCommands()

			command, ok := availableCommands[commandName]

			if !ok {
				fmt.Println("Invalid command")
				continue
			}
			
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
	
		}
}



type cliCommand struct {

	name string
	description string
	callback func(*config, ...string) error

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
		"explore": {
			name: "explore {location_area}",
			description: "Lists the pokemon in a location area",
			callback: callbackExplore,
		},
		"catch": {
			name: "catch {pokemon_name}",
			description: "Attempts to catch a Pokemon",
			callback: callbackCatch,
		},
		"inspect": {
			name: "inspect {pokemon_name}",
			description: "Prints information about a caught Pokemon",
			callback: callbackInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Prints all caught Pokemon",
			callback: callbackPokedex,
		},


	}
}


func cleanInput(str string) []string {

	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
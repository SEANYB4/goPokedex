package main



import (
	"errors"
	"fmt"

)


func callbackInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("You haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {

		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	for _, typ := range pokemon.Types {
		fmt.Printf(" - types: %s", typ.Type.Name)
	}

	return nil

}
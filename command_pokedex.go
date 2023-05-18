package main


import (
	"fmt"

)

func callbackPokedex(cfg *config, args ...string) error {

	fmt.Println("Pokemon in Pokedex:")

	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("No Pokemon caught yet...")
		return nil
	}

	for _, pokemon := range cfg.caughtPokemon {

		fmt.Println(pokemon.Name)
	}


	return nil

}
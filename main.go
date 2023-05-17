package main

import (
	"github.com/SEANYB4/goPokedex/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string

}

func main() {

	cfg := config {
		pokeapiClient : pokeapi.NewClient(time.Hour),
	}
	startRepl(&cfg)
	
}
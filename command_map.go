package main


import (
	
	"fmt"
	"log"
	"errors"
) 

func callbackMap(cfg *config) error {

	

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)

	if err != nil {
		return err
	}

	fmt.Println("Location areas:")

	for _, area := range resp.Results {


		fmt.Printf(" - %s", area.Name)
		fmt.Println("")

	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
	
}

func callbackMapB(cfg *config) error {

	if cfg.prevLocationAreaURL == nil {
		return errors.New("No previous page of locations to load")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)

	if err != nil {
		return err
	}

	fmt.Println("Location areas:")

	for _, area := range resp.Results {


		fmt.Printf(" - %s", area.Name)
		fmt.Println("")

	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
	
}
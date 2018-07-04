package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
	DexNo   int       `json:"id"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A PokemonSpecies struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func listCompleteDex() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println("==================")
	fmt.Println("Region:", strings.Title(responseObject.Name))
	fmt.Println("==================")

	for i := 0; i < len(responseObject.Pokemon); i++ {

		fmt.Println("Pokédex Entry: ", responseObject.Pokemon[i].EntryNo)
		fmt.Println("Pokémon Name: ", strings.Title(responseObject.Pokemon[i].Species.Name))
		fmt.Println("URL Reference: ", responseObject.Pokemon[i].Species.URL)
		fmt.Println("")
	}

	fmt.Println("==================")
}

func findPokemonByName() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokemon/mewtwo")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println("==================")
	fmt.Println("Pokedex Entry: ", responseObject.DexNo)
	fmt.Println("Pokemon Name: ", strings.Title(responseObject.Name))
	fmt.Println("==================")
	fmt.Println("")

}

func findPokemonByNumber() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokemon/1")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println("==================")
	fmt.Println("Pokedex Entry: ", responseObject.DexNo)
	fmt.Println("Pokemon Name: ", strings.Title(responseObject.Name))
	fmt.Println("==================")
	fmt.Println("")
}

func main() {
	var option string

	fmt.Print("\n==================")
	fmt.Print("\nMatt's Golang Pokedex")
	fmt.Print("\n==================\n")
	fmt.Print("\nSelect from the followng options:")
	fmt.Print("\n1. List complete Kanto Pokedex")
	fmt.Print("\n2. Find Pokemon stats by Name")
	fmt.Print("\n3. Find Pokemon stats by Number\n\n")

	fmt.Scan(&option)

	if option == "1" {
		listCompleteDex()
	} else if option == "2" {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Pokemon Name: ")
		r.ReadString('\n')
		findPokemonByName()
	} else if option == "3" {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Pokedex Entry: ")
		r.ReadString('\n')
		findPokemonByNumber()
	} else {
		fmt.Println("Not a valid entry, try again")
	}

}

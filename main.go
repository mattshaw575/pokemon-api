package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
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
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Pokemon Name: ")
	pokemonName, _ := r.ReadString('\n')

	u, err := url.Parse("http://pokeapi.co/api/v2/pokemon/")
	u.Path = path.Join(u.Path, pokemonName)
	URL := u.String()
	newURL := strings.ToLower(strings.TrimSuffix(URL, "%0A"))
	// println(newURL)

	response, err := http.Get(newURL)

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

	fmt.Println("\n==================")
	fmt.Println("Pokedex Entry: ", responseObject.DexNo)
	fmt.Println("Pokemon Name: ", strings.Title(responseObject.Name))
	fmt.Println("==================")
	fmt.Println("")
}

func findPokemonByNumber() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Pokedex Entry: ")
	pokemonNo, _ := r.ReadString('\n')

	var i int
	if (i >= 0) && (i <= 807) {
		fmt.Println("Searching for Pokemon...")
	} else {
		fmt.Println("Oak's words echoed... The National Pokedex only goes so far...")
		fmt.Println("Generations 1-7 have a range from #1 to #807")
		return
	}

	u, err := url.Parse("http://pokeapi.co/api/v2/pokemon/")
	u.Path = path.Join(u.Path, pokemonNo)
	URL := u.String()
	newURL := strings.TrimSuffix(URL, "%0A")
	// println(newURL)

	response, err := http.Get(newURL)

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

	fmt.Println("\n==================")
	fmt.Println("Pokedex Entry: ", responseObject.DexNo)
	fmt.Println("Pokemon Name: ", strings.Title(responseObject.Name))
	fmt.Println("==================")
	fmt.Println("")
}

func main() {
	var option string

	fmt.Println("==================")
	fmt.Println("Matt's Golang Pokedex")
	fmt.Println("==================\n")
	fmt.Println("Select from the followng options:")
	fmt.Println("1. List complete Kanto Pokedex")
	fmt.Println("2. Find Pokemon stats by Name")
	fmt.Println("3. Find Pokemon stats by Number\n")

	fmt.Scan(&option)

	if option == "1" {
		listCompleteDex()

	} else if option == "2" {
		findPokemonByName()

	} else if option == "3" {
		findPokemonByNumber()

	} else {
		fmt.Println("Not a valid entry, try again")
	}

}

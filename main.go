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
	"time"
)

// A Response struct to map the Entire Response
type Response struct {
	Name        string    `json:"name"`
	Pokemon     []Pokemon `json:"pokemon_entries"`
	DexNo       int       `json:"id"`
	PokemonType []Type    `json:"types"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A Type will list (up to) both types a given Pokemon can have
type Type struct {
	TypeSlots int         `json:"slot"`
	Type      PokemonType `json:"type"`
}

// A PokemonSpecies struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// A PokemonType struct to map our Pokemon's Species which includes it's name
type PokemonType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func listCompleteDexByRegion() {
	r := bufio.NewReader(os.Stdin)

	fmt.Println("\nWhich region do you want the pokemon listings from?")
	fmt.Println("1. National (across all generations)")
	fmt.Println("2. Kanto (gen 1)")
	fmt.Println("3. Johto (gen 2)")
	fmt.Println("4. Hoenn (gen 3)")
	fmt.Println("5. Sinnoh (gen 4)")
	fmt.Println("6. Unova (gen 5)")
	fmt.Println("7. Kalos (gen 6)")
	fmt.Println("8. Return")
	fmt.Println()

	pokedexRegion, _ := r.ReadString('\n')
	u, err := url.Parse("http://pokeapi.co/api/v2/pokedex/")

	switch {
	case strings.TrimSpace(pokedexRegion) == "6":
		pokedexRegion = "8"

	case strings.TrimSpace(pokedexRegion) == "7":
		pokedexRegion = "12"

	case strings.TrimSpace(pokedexRegion) == "8":
		fmt.Println("\nJumping back to main function...")
		time.Sleep(2000 * time.Millisecond)
		main()
	}

	u.Path = path.Join(u.Path, pokedexRegion)
	URL := u.String()
	dexURL := strings.ToLower(strings.TrimSuffix(URL, "%0A"))

	fmt.Println("\nProcessing request...")
	time.Sleep(2000 * time.Millisecond)
	response, err := http.Get(dexURL)

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

	fmt.Println("======================")
	fmt.Println("Region:", strings.Title(responseObject.Name))
	fmt.Println("=====================")

	for i := 0; i < len(responseObject.Pokemon); i++ {

		fmt.Println("\nPokédex Entry: ", responseObject.Pokemon[i].EntryNo)
		fmt.Println("Pokémon Name: ", strings.Title(responseObject.Pokemon[i].Species.Name))
		// fmt.Println("Pokemon Type: ", strings.Title(responseObject.PokemonType[i].Type.Name))
		fmt.Println("URL Reference: ", responseObject.Pokemon[i].Species.URL)
		fmt.Println()
	}
	fmt.Println("=====================")
	fmt.Println()
}

func findPokemonByNameOrNumber() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Pokemon Name or Number: ")
	pokemonNameOrNumber, _ := r.ReadString('\n')

	u, err := url.Parse("http://pokeapi.co/api/v2/pokemon/")
	u.Path = path.Join(u.Path, pokemonNameOrNumber)
	URL := u.String()
	newURL := strings.ToLower(strings.TrimSuffix(URL, "%0A"))

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

	if responseObject.DexNo == 0 {
		fmt.Println("\nOak's words echoed... The National Pokedex only goes so far...")
		fmt.Println("Generations 1-7 have a range from #1 to #721")
		time.Sleep(5000 * time.Millisecond)
		main()
	}

	fmt.Println("\n=====================")
	fmt.Println("Pokedex Entry: ", responseObject.DexNo)
	fmt.Println("Pokemon Name: ", strings.Title(responseObject.Name))

	for i := 0; i < len(responseObject.PokemonType); i++ {
		fmt.Println("Pokemon Type: ", strings.Title(responseObject.PokemonType[i].Type.Name))
	}

	fmt.Println("=====================")
	fmt.Println()
}

func additionalSearches() {
	var retry string

	fmt.Println("Continue Operations? ")
	fmt.Scan(&retry)

	switch {
	case strings.HasPrefix(retry, "y"):
		fmt.Println("\nJumping back to main function...")
		time.Sleep(2000 * time.Millisecond)
		main()

	case strings.HasPrefix(retry, "Y"):
		fmt.Println("\nJumping back to main function...")
		time.Sleep(2000 * time.Millisecond)
		main()

	default:
		fmt.Println("\nExiting operations")
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	var option int

	print("\033[H\033[2J") // Clear the terminal when running
	fmt.Println("=====================")
	fmt.Println("Matt's Golang Pokedex")
	fmt.Println("=====================")
	fmt.Println("\nSelect from the followng options:")
	fmt.Println("1. List all Pokemon by selected Pokedex")
	fmt.Println("2. Find Pokemon stats by Name or Number")
	fmt.Println()

	fmt.Scan(&option)

	switch {
	case option == 1:
		listCompleteDexByRegion()
		additionalSearches()

	case option == 2:
		findPokemonByNameOrNumber()
		additionalSearches()

	default:
		fmt.Println("Not a valid entry, try again")
		time.Sleep(2000 * time.Millisecond)
		main()
	}

}

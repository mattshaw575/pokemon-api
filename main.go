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

func listCompleteDexByRegion() {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("\nWhich region do you want the pokemon listings from?")
	fmt.Println("1. National (across all generations)")
	fmt.Println("2. Kanto (gen 1)")
	fmt.Println("3. Johto (gen 2)")
	fmt.Println("4. Hoenn (gen 3)")
	fmt.Println("5. Sinnoh (gen 4)") // 5+6
	fmt.Println("6. Unova (gen 5)")  // 8+9
	fmt.Println("7. Kalos (gen 6)")  //12+13+14
	fmt.Println()

	pokedexRegion, _ := r.ReadString('\n')

	u, err := url.Parse("http://pokeapi.co/api/v2/pokedex/")
	u.Path = path.Join(u.Path, pokedexRegion)
	URL := u.String()
	newURL := strings.ToLower(strings.TrimSuffix(URL, "%0A"))

	response, err := http.Get(newURL)

	// response, err := http.Get("http://pokeapi.co/api/v2//kanto/")
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
		fmt.Println()
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
	fmt.Println()
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
		fmt.Println("Generations 1-7 have a range from #1 to #721")
		return
	}

	u, err := url.Parse("http://pokeapi.co/api/v2/pokemon/")
	u.Path = path.Join(u.Path, pokemonNo)
	URL := u.String()
	newURL := strings.TrimSuffix(URL, "%0A")

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
	fmt.Println()
}

func additionalSearches() {
	var retry string

	fmt.Println("==================")
	fmt.Println("Continue Operations? ")

	fmt.Scan(&retry)

	// if retry == ("Yes") {
	// 	listCompleteDexByRegion()

	// } else {
	// 	fmt.Println("Exiting operations")
	// }

	switch {
	case retry == "Yes":
		main()

	case retry == "yes":
		main()

	case retry == "Y":
		main()

	case retry == "y":
		main()

	default:
		fmt.Println("Exiting operations")
	}
}

func main() {
	var option string

	print("\033[H\033[2J") // Clear the terminal when running
	fmt.Println("==================")
	fmt.Println("Matt's Golang Pokedex")
	fmt.Println("==================")
	fmt.Println("\nSelect from the followng options:")
	fmt.Println("1. List complete Kanto Pokedex")
	fmt.Println("2. Find Pokemon stats by Name")
	fmt.Println("3. Find Pokemon stats by Number")
	fmt.Println()

	fmt.Scan(&option)

	if option == "1" {
		listCompleteDexByRegion()
		additionalSearches()

	} else if option == "2" {
		findPokemonByName()
		additionalSearches()

	} else if option == "3" {
		findPokemonByNumber()
		additionalSearches()

	} else {
		fmt.Println("Not a valid entry, try again")
	}

}

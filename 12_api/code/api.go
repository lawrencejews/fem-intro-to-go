package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

//Create a route for people, write a function that return sucessful

//Planet is a planet type
type Planet struct{
	Name string `json: "name"`
	Population string `json: "population"`
	Terrain string `json: "terrain"`
}

//Person is a person type
type Person struct {
	Name string	`json: "name"`
	HomeWorldURL string	`json: "homeworld"`
	HomeWorld Planet 
}

//AllPeople is a collection of Person types
type AllPeople struct {
	People []Person `json: "results"`
}

//
func (p *Person) getHomeworld(){
	res, err := http.Get(p.HomeWorldURL)
	if err != nil {
		log.Print("Error fetching homeworld", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil{
		log.Print("Error reading response body", err)

		json.Unmarshal(bytes, &p.HomeWorld)
	}
}

func getPeople(w http.ResponseWriter, r *http.Request){
	res, err := http.Get(BaseURL + "people")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star war names")
		
	}

	fmt.Println(res)

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parser request body")
	}

	var people AllPeople

	fmt.Println(string(bytes))

	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(people)

	for _, pers := range people.People{
		pers.getHomeworld()
		fmt.Println(pers)
	}

}

func main() {

	http.HandleFunc("/people", getPeople)
	fmt.Println("Serving on: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

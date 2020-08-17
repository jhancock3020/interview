package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/mux"
)

// PROBLEM DESCRIPTION:
// Your goal here is to design an API that allows for hero tracking, much like the Vue problem
// You are to implement an API (for which the skeleton already exists) that has the following capabilities
// - Get      : return a JSON representation of the hero with the name supplied
// - Make     : create a superhero according to the JSON body supplied
// - Calamity : a calamity of the supplied level requires heroes with an equivalent or greater combined powerlevel to address it.
//              Takes a calamity with powerlevel and at least 1 hero. On success return a 200 with json response indicating the calamity has been resolved.
//              Otherwise return a response indicating that the heroes were not up to the task. Addressing a calamity adds 1 point of exhaustion.
// - Rest     : recover 1 point of exhaustion
// - Retire   : retire a superhero, someone may take up his name for the future passing on the title
// - Kill     : a superhero has passed away, his name may not be taken up again.

// On success all endpoints should return a status code 200.

// If a hero reaches an exhaustion level of maxExhaustion then they die.

// You are free to decide what your API endpoints should be called and what shape they should take. You can modify any code in this file however you'd like.

// NOTE: you may want to install postman or another request generating software of your choosing to make testing easier. (api is running on localhost port 8081)

// NOTE the second: the API is receiving asynchronous requests to manage our super friends. As such, your hero access should be thread-safe.
// Even if the operations are extremely lightweight we want to make our application scalable.
// Your multithreaded protection should allow the API to still be performant even if these requests took a reasonably long time. This means that
// a global lock/mutex that makes the API only handle 1 request at a time is not an acceptable solution to this problem

// NOTE the third: There are many ways to make whatever package-level tracking you implement thread-safe, feel free to change anything about this file (without changing the functionality of the program) to do so.
// i.e. add package-level maps, add functions that take the hero struct as reference, modify the hero struct, creating access control paradigms
// I highly recommend looking into channels, mutexes, and other golang memory management patterns and pick whatever you're most comfortable with.
// For mad props: a timeout on the memory management process which returns a resource not available.

// Bonus: If you're having fun (this is by no means necessary) you can make the calamity hold the heroes up for a time and delay their unlocking in a go-routine
// example:
// go func(h *hero) {
//     time.Sleep(calamityTime)
//     // release lock on hero
// }(heroPtr)
//Webpages used as a reference: https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da
//https://gist.github.com/aodin/9493190
//https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
//https://gobyexample.com/string-formatting
//https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
//https://blog.golang.org/json
//https://yourbasic.org/golang/json-example/
//https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e#:~:text=When%20the%20HTTP%20server%20is,method%20has%20the%20following%20signature.

var maxExhaustion = 10

type hero struct {
	PowerLevel int    `json:"PowerLevel"`
	Exhaustion int    `json:"Exhaustion"`
	Name       string `json:"Name"`
	mu 		   sync.Mutex

}
type calamity struct {
	PowerLevel int          `json:"PowerLevel"`
	HeroesInvolved []string `json:"heroesInvolved"`
	muCalam 		   sync.Mutex
}

// TODO: add storage and memory management
//Slice containing all alive hero structs
var allHeroes []hero
//Slice containing all dead hero structs
var deadHeroes []hero
// TODO: add more routines

// - Get      : return a JSON representation of the hero with the name supplied
func heroGet(w http.ResponseWriter, r *http.Request) {
	// TODO: access the global tracking to return the hero object
	var name string
	var ok bool
	// Vars checks name from URL. If handle not okay,
	// returns that value is not valid in console
	if name, ok = mux.Vars(r)["name"]; !ok {
		// TODO: Handle not ok
		fmt.Println("value:", name, "is ok:", ok)
		return
	}
	//for loop goes through allHeroes slice, returns index (we don't care about index)
	// and hero var
	for _, singleEvent := range allHeroes {
		// if hero var name is equal to previously found name value
		if singleEvent.Name == name {
			//Show hero struct on page as JSON
			json.NewEncoder(w).Encode(singleEvent)
			// Returns a status code 200
			//w.WriteHeader(http.StatusOK)
		}
	}
}
// - Make     : create a superhero according to the JSON body supplied
func heroMake(w http.ResponseWriter, r *http.Request) {
	// TODO : create a hero and add to some sort of global tracking
	// Creates a new hero var called newHero
	var newHero hero
	// Reads in JSON info
	content, err := ioutil.ReadAll(r.Body)
	// If there is an error with JSON data, writes string to user
	if err != nil {
		fmt.Fprintf(w, "error with JSON format")
		return
	}
	// Call json.Unmarshal to decode,
	// passing it a []byte of JSON data and a pointer to newHero variable
	json.Unmarshal(content, &newHero)
	for _, currHero := range allHeroes {
		// If currHero name is equal to newHero name
		if currHero.Name == newHero.Name {
			// Write to page
			fmt.Fprintf(w, "The hero name %v is already in use", newHero.Name)
			return
		}
	}
	// For loop goes through deadHeroes
	for _, deadHero := range deadHeroes {
		// If deadHero name is equal to newHero name
		if deadHero.Name == newHero.Name {
			// Write to page
			fmt.Fprintf(w, "The hero name %v has been retired", newHero.Name)
			return
		}
	}
	// Add now populated newHero var to newHero slice
	allHeroes = append(allHeroes, newHero)
	// prints JSON content to console
	fmt.Println(string(content))
	// Returns a status code 200
	w.WriteHeader(http.StatusOK)
}
// - Kill     : a superhero has passed away, his name may not be taken up again.
func heroKill(w http.ResponseWriter, r *http.Request) {
	var name string
	var ok bool
	// Vars checks name from URL. If handle not okay,
	// returns that value is not valid in console
	if name, ok = mux.Vars(r)["name"]; !ok {
		// TODO: Handle not ok
		fmt.Println("value:", name, "is ok:", ok)
		return
	}
	// Goes through allHeroes slice
	for i, currHero := range allHeroes {
		// If currHero name is equal to URL name
		if currHero.Name == name {
			// Add allHeroes index to deadHeroes slice
			deadHeroes = append(deadHeroes, allHeroes[i])
			// Shift allHeroes elements at the right of deleted element to the left
			allHeroes = append(allHeroes[:i], allHeroes[i+1:]...)
			// Write to user page
			fmt.Fprintf(w, "The hero with the name %v has been killed", name)
		}
	}
}
// - Retire   : retire a superhero, someone may take up his name for the future passing on the title
func heroRetire(w http.ResponseWriter, r *http.Request) {
	var name string
	var ok bool
	// Vars checks name from URL. If handle not okay,
	// returns that value is not valid in console
	if name, ok = mux.Vars(r)["name"]; !ok {
		// TODO: Handle not ok
		fmt.Println("value:", name, "is ok:", ok)
		return
	}
	// Goes through allHeroes slice
	for i, currHero := range allHeroes {
		// If currHero name is equal to URL name
		if currHero.Name == name {
			// Shift allHeroes elements at the right of deleted element to the left
			allHeroes = append(allHeroes[:i], allHeroes[i+1:]...)
			// Write to user page
			fmt.Fprintf(w, "The hero with the name %v has retired", name)
		}
	}
}
// - Rest     : recover 1 point of exhaustion
func heroRest(w http.ResponseWriter, r *http.Request) {
	var name string
	var ok bool
	// Vars checks name from URL. If handle not okay,
	// returns that value is not valid in console
	if name, ok = mux.Vars(r)["name"]; !ok {
		// TODO: Handle not ok
		fmt.Println("value:", name, "is ok:", ok)
		return
	}
	// Goes through allHeroes slice
	for i, currHero := range allHeroes {
		// If currHero name is equal to URL name
		// and currHero Exhaustion is greater than zero
		if currHero.Name == name && currHero.Exhaustion > 0{
			// Decrease currHero Exhaustion by one
			currHero.Exhaustion -= 1
			// Add change to allHeroes slice by
			// replaced the indexed element with currHero
			allHeroes[i] = currHero
			fmt.Fprintf(w, "hero list %v \n", allHeroes)
			// Write updated currHero to user page as JSON
			json.NewEncoder(w).Encode(currHero)
		}
	}
}
// - Calamity : a calamity of the supplied level requires heroes with an equivalent or greater combined powerlevel to address it.
//              Takes a calamity with powerlevel and at least 1 hero. On success return a 200 with json response indicating the calamity has been resolved.
//              Otherwise return a response indicating that the heroes were not up to the task. Addressing a calamity adds 1 point of exhaustion.
func heroCalamity(w http.ResponseWriter, r *http.Request) {
	// Creates a new calamity struct called newCalamity
	var newCalamity calamity
	// Creates total combined hero power and starts at zero
	var totalPower int  = 0
	// Reads in JSON info
	content, err := ioutil.ReadAll(r.Body)
	// If there is an error with JSON data
	if err != nil {
		// Writes to user page and ends func
		fmt.Fprintf(w, "error with JSON format")
		return
	}
	// Call json.Unmarshal to decode,
	// passing it a []byte of JSON data and a pointer to newHero variable
	json.Unmarshal(content, &newCalamity)
	// Goes through HeroesInvolved slice
	for _, calamityHero := range newCalamity.HeroesInvolved {
		// Goes through allHeroes slice
		for i, currHero := range allHeroes {
			// If currHero name is equal to calamityHero
			if currHero.Name == calamityHero{
				// currHero PowerLevel is added to totalPower
				totalPower += currHero.PowerLevel
				// currHero Exhaustion is increased by one
				currHero.Exhaustion += 1
				// Add change to allHeroes slice by
				// replaced the indexed element with currHero
				allHeroes[i] = currHero
				// If currHero Exhaustion is equal to maxExhaustion
				if currHero.Exhaustion == maxExhaustion{
					// Add allHeroes index to deadHeroes slice
					deadHeroes = append(deadHeroes, allHeroes[i])
					//Shift allHeroes elements at the right of deleted element to the left
					allHeroes = append(allHeroes[:i], allHeroes[i+1:]...)
					// Write to user page
					fmt.Fprintf(w, "The hero with the name %v has been killed \n", currHero.Name)
				}
			}
		}
	}
	// If Calamity PowerLevel is less than totalPower of all heroes
	if newCalamity.PowerLevel <= totalPower{
		// Write to user page
		fmt.Fprintf(w, "The calamity with powerlevel %d has been handled by heroes with powerlevel of %d consisting of %v!", newCalamity.PowerLevel, totalPower, newCalamity.HeroesInvolved)
		// Returns a status code 200
		w.WriteHeader(http.StatusOK)
	} else{
		// Write to user page
		fmt.Fprintf(w, "The calamity with powerlevel %d has NOT been handled by heroes with powerlevel of %d consisting of %v!", newCalamity.PowerLevel, totalPower, newCalamity.HeroesInvolved)
	}
}
func linkRoutes(r *mux.Router) {
	r.HandleFunc("/hero", heroMake).Methods("POST")
	r.HandleFunc("/hero/{name}", heroGet).Methods("GET")
	// TODO: add more routes
	r.HandleFunc("/hero/kill_{name}", heroKill).Methods("DELETE")
	r.HandleFunc("/hero/retire_{name}", heroRetire).Methods("DELETE")
	r.HandleFunc("/hero/rest_{name}", heroRest).Methods("PATCH")
	r.HandleFunc("/calamity", heroCalamity).Methods("POST")
}
func main() {
	// create a router
	router := mux.NewRouter()
	// and a server to listen on port 8081
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8081),
		Handler: router,
	}
	// link the supplied routes
	linkRoutes(router)
	// wait for requests
	log.Fatal(server.ListenAndServe())
}
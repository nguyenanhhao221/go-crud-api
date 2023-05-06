package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/mux"

	// "math/rand"
	"net/http"
	// "strconv"
)

/*
struct is like an object in javascript
the backtick “ is a value pair which act as additional information.
For example in this case, the ID key when encoding or decoding as json should be in key "id" instead of "ID"
*/

type Movie struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	/*Using a pointer type for the Director field allows for better memory management and can improve performance.
	If the Director field was declared as a non-pointer type, then a copy of the entire Director struct would need to be created every time a new Movie instance is created.
	This can become expensive if there are many Movie instances or if the Director struct is large.
	*/
	Director *Director
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

// getMovies send the list of movies as json to client
func getMovies(w http.ResponseWriter, r *http.Request) {
	// Set the response header to let client know that it a json
	w.Header().Set("Content-Type", "application/json")

	// Need to use json package to encode the movies slice into json format
	// These syntax is check if the encoding process have any error
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		log.Printf("Error encoding movies: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
func main() {
	// Initialize a new router
	r := mux.NewRouter()

	// Populated the movies slice, a tricky way
	movies = append(movies, Movie{ID: "1", Isbn: "123", Title: "Movie one", Director: &Director{FirstName: "James", LastName: "Gunn"}})
	movies = append(movies, Movie{ID: "2", Isbn: "289137", Title: "Movie Two", Director: &Director{FirstName: "Nguyen", LastName: "Duy"}})

	// Handle for the route /movies
	r.HandleFunc("/movies", getMovies).Methods("GET")

	// Start the app
	port := "8080"
	fmt.Printf("Starting server at port:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

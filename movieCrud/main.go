package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorrila/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&Movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie() {

	w.Header().Set("Content-Type", "applicaiton/json")
	//params
	params := mux.Vars(r)
	//loop over movies, range
	//delete the movie with id that you have sent
	//add a new movie

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var mvoie Movie
		}
	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie one", Director: &Director{Firstname: "Samra", Lastname: "azizi"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie two", Director: &Director{Firstname: "fahime", Lastname: "ahmadi"}})
	r.HandleFux("/movies", getMovie).Methods("GET")
	r.handleFunc("/movies/{id}", getMovie).Methods("GET")
	r.handleFunc("/movies", createMovie).Methods("POST")
	r.handleFunc("movies/{id}", updateMovie).Methods("PUT")
	r.handleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

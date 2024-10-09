package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id:`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}


func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index, val := range movies {
		if val.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, val := range movies{
		if val.ID == params["id"]{
			json.NewEncoder(w).Encode(val)
			return
		}
	}
}


func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	 _ = json.NewDecoder(r.Body).Decode(&movie)
	 movie.ID = strconv.Itoa(rand.Intn(1000000000))
	 movies = append(movies, movie)
	 json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, val := range movies{
		if val.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	 movie.ID = params["id"]
	 movies = append(movies, movie)
	 json.NewEncoder(w).Encode(movie)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1",
		Isbn: "243435",
		Title: "Life of Pi",
		Director: &Director{
			FirstName: "Jay",
			LastName: "Pawar",
		},
	})

	movies = append(movies, Movie{
		ID: "2",
		Isbn: "4646468",
		Title: "Martian",
		Director: &Director{
			FirstName: "John",
			LastName: "Doe",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", ISBN: "79878978", Title: "one", Director: &Director{Firstname: "a", Lastname: "waye"}})
	movies = append(movies, Movie{ID: "2", ISBN: "12345678", Title: "two", Director: &Director{Firstname: "b", Lastname: "smith"}})
	movies = append(movies, Movie{ID: "3", ISBN: "78901234", Title: "three", Director: &Director{Firstname: "c", Lastname: "johnson"}})
	movies = append(movies, Movie{ID: "4", ISBN: "34567890", Title: "four", Director: &Director{Firstname: "d", Lastname: "lee"}})

	r.HandleFunc("/movie", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(movies)
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	//set json content
	writer.Header().Set("Content-Type", "application/json")
	//param
	params := mux.Vars(request)
	//loop over the movie
	//delete the movie with the id
	//add a new movie
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = params["id"]
			json.NewEncoder(writer).Encode(movie)
			return
		}
	}
}

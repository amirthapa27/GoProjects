package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// creating a fake DB
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

//get all movies

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //setting content type as json
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

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)  //will decode the recieved data and store it in the movie variable
	movie.ID = strconv.Itoa(rand.Intn(1000000)) //converting integer ID to string
	movies = append(movies, movie)              //apending the newly created movie

	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //delete the existing data in that id

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {
	r := mux.NewRouter() //creating a new router

	//appending data in the fake db
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie one", Director: &Director{FirstName: "Amir", LastName: "Thapa"}})
	movies = append(movies, Movie{ID: "2", Isbn: "5496821", Title: "Movie Two", Director: &Director{FirstName: "Daksh", LastName: "Shirodkar"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")          //handler
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")       //handler
	r.HandleFunc("/movie", createMovie).Methods("POST")        //handler
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")    //handler
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE") //handler

	fmt.Println("Starting port 8000")
	log.Fatal(http.ListenAndServe(":8000", r)) //starting the server at port 8000
}

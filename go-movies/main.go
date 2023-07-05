package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstname"`
	Lastname string `josn:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json") //still dont understand this. all i know is that it encode from normal string to json type 
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies{

		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for _, item :=range movies{

		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie 
	_ = json.NewDecoder(r.Body).Decode(&movie) //i dont understand why need to decode
	movie.ID = strconv.Itoa(rand.Intn(1000000000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovies(w http.ResponseWriter, r *http.Request){

	//set json content type
	w.Header().Set("Content-Type","application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies, range
	//delete the movie with the i.d that have been chosen
	//add a new movie - get params by postman

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie 
			_ = json.NewDecoder(r.Body).Decode(&movie) // i dont understand why decode r.body
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie) // i dont understand why encode w..
			return
		}
	}

}



func main(){
	r := mux.NewRouter()


	movies = append(movies, Movie{ID: "090911", Isbn: "234443", Title: "one town", Director : &Director{FirstName:"Ahmad", Lastname: "Rasyid"}})
	movies = append(movies, Movie{ID: "988821", Isbn: "766543", Title: "two town", Director : &Director{FirstName: "Sukida", Lastname: "Water"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
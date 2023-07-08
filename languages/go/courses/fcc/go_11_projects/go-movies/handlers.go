package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")

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

	_ = json.NewDecoder(r.Body).Decode(&movie) // TODO: Bad error handling!
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set content type
	w.Header().Set("Content-Type", "application/json")

	// read params
	params := mux.Vars(r)

	// range over movies and delete the specified movie, add the replacement using the same id
	for idx, itm := range movies {
		if itm.ID == params["id"] {
			// found our mark
			movies = append(movies[:idx], movies[idx+1:]...) // append up to our index, them beyond it too, removing the indexed movie
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // TODO: Bad error handling
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}

}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/StalkR/imdb"
	"github.com/gorilla/mux"
	"github.com/mmorejon/microservices-docker-go-mongodb/movies/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	ret_movies := []models.Movie{}
	client := http.DefaultClient
	results, err := imdb.SearchTitle(client, "matrix")
	app.infoLog.Println("Movies have been listed from Imdb")
	for _, value := range results {
		rating := 0.0
		if rating, err := strconv.ParseFloat(value.Rating, 32); err == nil {
			fmt.Printf("%T, %v\n", rating, rating)
		}
		objID, err := primitive.ObjectIDFromHex(value.ID)
		if err != nil {
			app.infoLog.Println(err)
		}
		var mov = models.Movie{
			ID:          objID,
			Title:       value.Name,
			Description: value.Description,
			ImdbUrl:     value.URL,
			Rating:      rating,
			CreatedOn:   time.Time{},
		}
		ret_movies = append(ret_movies, mov)
		//app.infoLog.Println(index, ":", value.Name)
		////app.infoLog.Println("Rating:", value.Rating, "Rated By #", value.RatingCount)
		//app.infoLog.Println("IMDB Url:", value.URL)
	}

	//movies, err := app.movies.All()
	//if err != nil {
	//	app.serverError(w, err)
	//}

	// Convert movie list into json encoding
	b, err := json.Marshal(ret_movies)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Movies have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	client := http.DefaultClient
	results, err := imdb.SearchTitle(client, "matrix")
	app.infoLog.Println("Movies have been listed from Imdb")
	for index, value := range results {
		app.infoLog.Println(index, ":", value.Name)
		//app.infoLog.Println("Rating:", value.Rating, "Rated By #", value.RatingCount)
		app.infoLog.Println("IMDB Url:", value.URL)
	}
	// Find movie by id
	m, err := app.movies.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Movie not found")
			return
		}
		// Any other error will send an internal server error
		app.serverError(w, err)
	}

	// Convert movie to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a movie")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	// Define movie model
	var m models.Movie
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	// Insert new movie
	insertResult, err := app.movies.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New movie have been created, id=%s", insertResult.InsertedID)
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete movie by id
	deleteResult, err := app.movies.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d movie(s)", deleteResult.DeletedCount)
}

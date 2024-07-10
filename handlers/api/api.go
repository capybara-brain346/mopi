package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/capybarabrain346/mopi/handlers/database"
	"github.com/capybarabrain346/mopi/models"
	"github.com/go-chi/chi/v5"
)

func GetMovie(w http.ResponseWriter, r *http.Request) {
	movie_name := chi.URLParam(r, "movie_name")
	dbResponse, err := database.DBGet(movie_name)

	if err != nil {
		log.Println(err)
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching movie info for %s", movie_name)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(Jsonify(dbResponse)))
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var M models.MovieSchema
	err := json.NewDecoder(r.Body).Decode(&M)

	if err != nil {
		fmt.Println("Error parsing json request.")
	}
	fmt.Println(M.Name)
	dbResponse, err := database.DBWrite(M)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Error entering movie: %s", M.Name)))
	}

	if dbResponse > 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Database updated with \n %s", Jsonify(M))))
	}

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	movie_name := chi.URLParam(r, "movie_name")
	dbResponse, err := database.DBDelete(movie_name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Error deleting movie: %s", movie_name)))
	}

	if dbResponse > 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Deleted movie : %s", movie_name)))
	}

}

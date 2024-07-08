package api

import (
	"fmt"
	"net/http"

	"github.com/capybarabrain346/mopi/handlers/database"
	"github.com/capybarabrain346/mopi/models"
	"github.com/go-chi/chi/v5"
)

func GetMovie(w http.ResponseWriter, r *http.Request) {
	movie_name := chi.URLParam(r, "movie_name")
	movie_info, err := database.DBGet(movie_name)

	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching movie info for %s", movie_name)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(JsonifyGetMovie(movie_info)))
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	movie_name := chi.URLParam(r, "movie_name")
	movie_release_date := chi.URLParam(r, "movie_release_date")

	new_movie_object := models.MovieSchema{
		Title:       movie_name,
		ReleaseDate: movie_release_date,
	}

	create_movie := database.DBWrite(new_movie_object)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(create_movie))
}

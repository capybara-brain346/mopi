package database

import (
	"errors"
	"fmt"

	"github.com/capybarabrain346/mopi/models"
)

func DBGet(movieName string) (models.MovieSchema, error) {
	var dbResponse models.MovieSchema

	for _, movie_info := range models.MovieDB {
		if movie_info.Title == movieName {
			dbResponse = movie_info
			return dbResponse, nil
		}
	}
	return dbResponse, errors.New("movie not found")
}

func DBWrite(create_movie models.MovieSchema) string {
	moviedb_lenght := len(models.MovieDB) + 1
	models.MovieDB[moviedb_lenght] = create_movie

	dbResponse := fmt.Sprintf("%s record created", create_movie.Title)

	return dbResponse
}

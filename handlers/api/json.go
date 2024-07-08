package api

import (
	"encoding/json"

	"github.com/capybarabrain346/mopi/models"
)

func JsonifyGetMovie(movie_info models.MovieSchema) string {
	var json_reponse = map[string]string{"Title": movie_info.Title, "ReleaseDate": movie_info.ReleaseDate}
	var json_movie_info, err = json.Marshal(json_reponse)

	if err != nil {
		return "failed to jsonify!"
	}

	return string(json_movie_info)
}

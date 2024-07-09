package api

import (
	"encoding/json"

	"github.com/capybarabrain346/mopi/models"
)

func Jsonify(movie_info models.MovieSchema) string {
	var json_movie_info, err = json.Marshal(movie_info)

	if err != nil {
		return "failed to jsonify!"
	}

	return string(json_movie_info)
}

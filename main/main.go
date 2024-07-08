package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/capybarabrain346/mopi/handlers/api"
	"github.com/capybarabrain346/mopi/handlers/constants"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	json_welcome_message, err := json.Marshal(constants.JsonWelcome)
	if err != nil {
		log.Fatalf("Error!")
	}

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(string(json_welcome_message)))
	})

	r.Get("/movie/{movie_name}", api.GetMovie)
	r.Post("/createmovie/{movie_name}-{movie_release_date}", api.CreateMovie)
	http.ListenAndServe(":3000", r)
}

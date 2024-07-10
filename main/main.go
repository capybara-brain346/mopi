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
	r.Post("/movie", api.CreateMovie)
	r.Delete("/movie/{movie_name}", api.DeleteMovie)

	const port string = ":3000"
	http.ListenAndServe(port, r)
}

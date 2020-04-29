package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.HandleFunc("/", HelloServer)

	err := http.ListenAndServe(":"+defaultPort, router)
	if err != nil {
		panic("failed to start server")
	}
	fmt.Println("server started")

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

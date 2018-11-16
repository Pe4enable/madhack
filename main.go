package main

import (
	. "./application/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	MadnessRouter("/rosbank", router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		Debug:            true,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))

}
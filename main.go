package main

import (
	"github.com/BankEx/madhack/router"
	"github.com/BankEx/madhack/config"
	"github.com/BankEx/madhack/services"
	"github.com/BankEx/madhack/repositories"
	"github.com/BankEx/madhack/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"flag"
)

func main() {
	configPathPtr := flag.String("config-path", "./config/config.yaml", "path to config.yaml")
	flag.Parse()

	// Config
	conf, err := config.Load(*configPathPtr)
	if err != nil {
		panic(err)
	}

	reader, err := services.New(conf)
	if err != nil {
		panic(err)
	}
	log.Printf("Rates reader is initialised")

	mongo, err := repositories.New(conf.Mongo)
	if err != nil {
		panic(err)
	}
	log.Printf("MongoWriter is initialised")

	handlers := handlers.New(mongo)
	r := router.New(handlers)

	log.Printf("Server is listening on %s port", conf.Port)
	log.Panic(http.ListenAndServe(conf.Port, r))

	select {}

}
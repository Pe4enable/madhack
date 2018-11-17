package main

import (
	. "./application/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	configPathPtr := flag.String("config-path", "./config/config.yaml", "path to config.yaml")
	flag.Parse()

	// Config
	conf, err := config.Load(*configPathPtr)
	if err != nil {
		panic(err)
	}

	cache := make(chan ratestates.RateState, conf.Cache)

	reader, err := services.New(conf)
	if err != nil {
		panic(err)
	}
	log.Printf("Rates reader is initialised")
	go reader.Start(cache)

	mongo, err := repositories.New(conf.Mongo)
	if err != nil {
		panic(err)
	}
	log.Printf("MongoWriter is initialised")
	go mongo.Start(cache)

	handlers := handlers.New(mongo)
	r := router.New(handlers)

	log.Printf("Server is listening on %s port", conf.Port)
	log.Panic(http.ListenAndServe(conf.Port, r))

	select {}

}
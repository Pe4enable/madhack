package handlers

import (
	"log"
	"net/http"
)

func (s *HandlersService) GetItem(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	result, err := s.repository.GetItem(name)

	if err != nil {
		log.Printf("Error during getting rate: %s.", err)
		jsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, result, http.StatusOK)
}

func (s *HandlersService) GetAllItems(w http.ResponseWriter, r *http.Request) {

	result, err := s.repository.GetAllItems()

	if err != nil {
		log.Printf("Error during getting rate: %s.", err)
		jsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, result, http.StatusOK)
}

func (s *HandlersService) HealthCheck(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, true, http.StatusOK)
}

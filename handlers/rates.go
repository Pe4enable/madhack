package handlers

import (
	"log"
	"net/http"
)

func (s *HandlersService) GetLastRate(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	date := r.URL.Query().Get("date")

	result, err := s.repository.GetRate(from, to, date)

	if err != nil {
		log.Printf("Error during getting rate: %s.", err)
		jsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, result, http.StatusOK)
}

func (s *HandlersService) GetLastAllRates(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	date := r.URL.Query().Get("date")

	result, err := s.repository.GetAllRates(from, date)

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

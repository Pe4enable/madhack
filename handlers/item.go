package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/BankEx/madhack/models"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"time"
)

func (s *HandlersService) AddItem(w http.ResponseWriter, r *http.Request) {
	item := models.Item{
		ID:      objectid.New(),
		Created: time.Now(),
	}

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		jsonErrorResponse(w, err, http.StatusBadRequest )
		return
	}

	err := s.repository.AddItem(item)

	if err != nil {
		log.Printf("Error during getting rate: %s.", err)
		jsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, nil, http.StatusOK)
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

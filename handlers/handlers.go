package handlers

import (
	"github.com/BankEx/madhack/repositories"
)

type HandlersService struct {
	repository *repositories.MongoRepository
}

func New(repository *repositories.MongoRepository) *HandlersService {
	return &HandlersService{repository: repository}
}

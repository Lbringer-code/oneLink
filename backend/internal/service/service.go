package service

import (
	"log/slog"

	"github.com/Lbringer-code/oneLink/backend/internal/repository"
)

type Service struct {
	repo *repository.Repository
	logger *slog.Logger
}

func New( repo *repository.Repository  , logger *slog.Logger) *Service {
	return &Service{
		repo : repo ,
		logger: logger,
	}
}
package handler

import (
	"log/slog"

	"github.com/Lbringer-code/oneLink/backend/internal/service"
)

type Handler struct {
	svc *service.Service
	logger *slog.Logger
	allowedOrigins []string
}

func New( svc *service.Service , logger *slog.Logger , allowedOrigins []string ) *Handler {
	return &Handler{
		svc : svc , 
		logger : logger , 
		allowedOrigins: allowedOrigins,
	}
}
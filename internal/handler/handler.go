package handler

import "github.com/lantonster/iam/internal/service"

type Handler struct {
	Auth *AuthHandler
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		Auth: newAuthHandler(service),
	}
}

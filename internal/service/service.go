package service

import "github.com/lantonster/iam/internal/repo"

type Service struct {
	Auth AuthService
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		Auth: newDefaultAuthService(repo),
	}
}

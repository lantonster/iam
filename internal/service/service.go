package service

import (
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/repo"
)

type Service interface {
	Auth() AuthService
}

type defaultService struct {
	auth AuthService
}

func NewDefaultService(conf *config.Config, repo repo.Repo) Service {
	return &defaultService{
		auth: newDefaultAuthService(conf, repo),
	}
}

func (s *defaultService) Auth() AuthService {
	return s.auth
}

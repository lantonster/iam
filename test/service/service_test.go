package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lantonster/iam/internal/service"
	"github.com/lantonster/iam/test/repo"
)

var (
	ctx *gin.Context
	srv *service.Service

	mockRepo     *repo.MockRepo
	mockUserRepo *repo.MockUserRepo
)

func setupMock(ctrl *gomock.Controller) {
	mockRepo = repo.NewMockRepo(ctrl)
	mockUserRepo = repo.NewMockUserRepo(ctrl)

	ctx = &gin.Context{}
	srv = service.NewService(mockRepo)
}

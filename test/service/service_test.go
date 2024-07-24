package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lantonster/iam/internal/model"
	"github.com/lantonster/iam/internal/service"
	"github.com/lantonster/iam/pkg/utils"
	"github.com/lantonster/iam/test/repo"
)

var (
	ctx *gin.Context
	srv *service.Service

	currentUser *model.User

	mockRepo     *repo.MockRepo
	mockUserRepo *repo.MockUserRepo
)

func setupMock(ctrl *gomock.Controller) {
	mockRepo = repo.NewMockRepo(ctrl)
	mockUserRepo = repo.NewMockUserRepo(ctrl)

	ctx = &gin.Context{}
	srv = service.NewService(mockRepo)

	// 在 context 中设置当前登陆用户的信息，模拟登陆校验中间的操作
	currentUser = &model.User{Id: 1, Username: "admin"}
	utils.SetUserIdToContext(ctx, currentUser.Id)
	utils.SetUsernameToContext(ctx, currentUser.Username)
}

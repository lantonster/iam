package service

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/model"
	"github.com/lantonster/iam/internal/service"
	"github.com/lantonster/iam/pkg/utils"
	"github.com/lantonster/iam/test/repo"
)

var (
	ctx *gin.Context
	srv service.Service

	currentUser *model.User

	mockRepo                 *repo.MockRepo
	mockUserRepo             *repo.MockUserRepo
	mockVerificationCodeRepo *repo.MockVerificationCodeRepo
)

func setupMock(ctrl *gomock.Controller) {
	conf := config.NewConfig()

	mockRepo = repo.NewMockRepo(ctrl)
	mockUserRepo = repo.NewMockUserRepo(ctrl)
	mockVerificationCodeRepo = repo.NewMockVerificationCodeRepo(ctrl)

	gin.SetMode(gin.ReleaseMode)
	w := httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	srv = service.NewDefaultService(conf, mockRepo)

	// 在 context 中设置当前登陆用户的信息，模拟登陆校验中间的操作
	currentUser = &model.User{Id: 1, Username: "admin"}
	utils.SetUserIdToContext(ctx, currentUser.Id)
	utils.SetUsernameToContext(ctx, currentUser.Username)
}

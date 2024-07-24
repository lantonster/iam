package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lantonster/cerrors"
	"github.com/lantonster/ecodes"
	"github.com/lantonster/iam/internal/dto"
	"github.com/lantonster/iam/internal/model"
	"github.com/lantonster/iam/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	setupMock(ctrl)

	password, salt := "123456", "654321"
	hashed, _ := utils.HashPassword(password, salt)
	foundUser := &model.User{Username: "existing_user", Password: hashed, Salt: salt}

	// 测试用户名不存在的情况
	req := &dto.AuthLoginRequest{Username: "nonexistent_user", Password: "password"}
	mockRepo.EXPECT().User().Return(mockUserRepo)
	mockUserRepo.EXPECT().GetUserByUsername(ctx, req.Username).Return(nil, cerrors.WithCode(ecodes.IAM_USERNAME_NOT_FOUND, "user not found"))
	res, err := srv.Auth.Login(ctx, req)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Equal(t, cerrors.Code(err), ecodes.IAM_USERNAME_NOT_FOUND)

	// 测试找到用户但密码错误的情况
	req = &dto.AuthLoginRequest{Username: foundUser.Username, Password: "wrong_password"}
	mockRepo.EXPECT().User().Return(mockUserRepo)
	mockUserRepo.EXPECT().GetUserByUsername(ctx, req.Username).Return(foundUser, nil)
	res, err = srv.Auth.Login(ctx, req)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Equal(t, cerrors.Code(err), ecodes.IAM_PASSWORD_ERROR)

	// 测试找到用户且密码正确的情况
	req = &dto.AuthLoginRequest{Username: foundUser.Username, Password: password}
	mockRepo.EXPECT().User().Return(mockUserRepo)
	mockUserRepo.EXPECT().GetUserByUsername(ctx, req.Username).Return(foundUser, nil)
	res, err = srv.Auth.Login(ctx, req)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

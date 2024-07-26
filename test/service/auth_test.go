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

	t.Run("测试用户名不存在的情况", func(t *testing.T) {
		mockRepo.EXPECT().User().Return(mockUserRepo)
		mockUserRepo.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(nil, cerrors.WithCode(ecodes.IAM_USERNAME_NOT_FOUND, "user not found"))

		req := &dto.AuthLoginRequest{Username: "nonexistent_user", Password: "password"}
		res, err := srv.Auth().Login(ctx, req)

		assert.Nil(t, res)
		assert.Error(t, err)
		assert.Equal(t, ecodes.IAM_USERNAME_NOT_FOUND, cerrors.Code(err))
	})

	t.Run("测试找到用户但密码错误的情况", func(t *testing.T) {
		mockRepo.EXPECT().User().Return(mockUserRepo)
		mockUserRepo.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(foundUser, nil)

		req := &dto.AuthLoginRequest{Username: foundUser.Username, Password: "wrong_password"}
		res, err := srv.Auth().Login(ctx, req)

		assert.Nil(t, res)
		assert.Error(t, err)
		assert.Equal(t, ecodes.IAM_PASSWORD_ERROR, cerrors.Code(err))
	})

	t.Run("测试找到用户且密码正确的情况", func(t *testing.T) {
		mockRepo.EXPECT().User().Return(mockUserRepo)
		mockUserRepo.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(foundUser, nil)

		req := &dto.AuthLoginRequest{Username: foundUser.Username, Password: password}
		res, err := srv.Auth().Login(ctx, req)

		assert.NotNil(t, res)
		assert.NoError(t, err)
	})
}

func TestAuthService_UserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	setupMock(ctrl)

	t.Run("测试已登陆的用户获取信息", func(t *testing.T) {
		res, err := srv.Auth().UserInfo(ctx)

		assert.NotNil(t, res)
		assert.NoError(t, err)
		assert.Equal(t, &dto.AuthUserInfoResponse{UserId: currentUser.Id, Username: currentUser.Username}, res)
	})
}

func TestAuthService_UsernameAvailable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	setupMock(ctrl)

	t.Run("测试用户名长度不合法的情况", func(t *testing.T) {
		invalid := []string{"", "1234567890123456789012345678901"}

		for _, username := range invalid {
			req := &dto.AuthUsernameAvailableRequest{Username: username}
			err := srv.Auth().UsernameAvailable(ctx, req)

			assert.Error(t, err)
			assert.Equal(t, ecodes.IAM_INVALID_USERNAME_LENGTH, cerrors.Code(err))
		}
	})

	t.Run("测试用户名字符不合法的情况", func(t *testing.T) {
		invalid := []string{" ", "$", "/t", "[]", "中文", "!", "@"}

		for _, username := range invalid {
			req := &dto.AuthUsernameAvailableRequest{Username: username}
			err := srv.Auth().UsernameAvailable(ctx, req)

			assert.Error(t, err)
			assert.Equal(t, ecodes.IAM_INVALID_USERNAME_FORMAT, cerrors.Code(err))
		}
	})

	t.Run("测试用户名重复的情况", func(t *testing.T) {
		mockRepo.EXPECT().User().Return(mockUserRepo)
		mockUserRepo.EXPECT().CheckUserNameDuplication(gomock.Any(), gomock.Any()).Return(true, nil)

		req := &dto.AuthUsernameAvailableRequest{Username: "exist"}
		err := srv.Auth().UsernameAvailable(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, ecodes.IAM_USERNAME_ALREADY_EXISTS, cerrors.Code(err))
	})

	t.Run("测试用户名可用的情况", func(t *testing.T) {
		mockRepo.EXPECT().User().Return(mockUserRepo)
		mockUserRepo.EXPECT().CheckUserNameDuplication(gomock.Any(), gomock.Any()).Return(false, nil)

		req := &dto.AuthUsernameAvailableRequest{Username: "non_exist"}
		err := srv.Auth().UsernameAvailable(ctx, req)

		assert.NoError(t, err)
	})
}

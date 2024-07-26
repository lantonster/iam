package handler

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lantonster/iam/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	setupMock(ctrl)

	mockService.EXPECT().Auth().Return(mockAuthService)
	mockAuthService.EXPECT().Login(gomock.Any(), gomock.Any()).Return(&dto.AuthLoginResponse{Token: "mocked_token"}, nil)

	t.Run("测试参数缺失的情况", func(t *testing.T) {
		w := doRequest("GET", "/auth/login?username=admin&password=", nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("测试参数拼错的情况", func(t *testing.T) {
		w := doRequest("GET", "/auth/login?Username=admin&Password=123456", nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("测试正确填入参数的情况", func(t *testing.T) {
		w := doRequest("GET", "/auth/login?username=admin&password=123456", nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestAuthHandler_UserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	setupMock(ctrl)

	t.Run("测试未登陆的情况", func(t *testing.T) {
		w := doRequestWithoutAuthorization("GET", "/auth/user_info", nil)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("测试登陆的情况", func(t *testing.T) {
		mockService.EXPECT().Auth().Return(mockAuthService)
		mockAuthService.EXPECT().UserInfo(gomock.Any()).Return(&dto.AuthUserInfoResponse{}, nil)

		w := doRequest("GET", "/auth/user_info", nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestAuthHandler_UsernameAvailable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	setupMock(ctrl)

	t.Run("测试参数缺失的情况", func(t *testing.T) {
		w := doRequest("GET", "/auth/username_available?username=", nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("测试参数拼错的情况", func(t *testing.T) {
		w := doRequest("GET", "/auth/username_available?Username=admin", nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("测试正确填入参数的情况", func(t *testing.T) {
		mockService.EXPECT().Auth().Return(mockAuthService)
		mockAuthService.EXPECT().UsernameAvailable(gomock.Any(), gomock.Any()).Return(nil)

		w := doRequest("GET", "/auth/username_available?username=admin", nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

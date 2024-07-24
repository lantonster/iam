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

	// 测试参数缺失的情况
	w := doRequest("GET", "/auth/login?username=admin&password=", nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// 测试参数拼错的情况
	w = doRequest("GET", "/auth/login?Username=admin&Password=123456", nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// 测试正确填入参数的情况
	w = doRequest("GET", "/auth/login?username=admin&password=123456", nil)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthHandler_UserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	setupMock(ctrl)

	// 测试未登陆的情况
	w := doRequestWithoutAuthorization("GET", "/auth/user_info", nil)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	mockService.EXPECT().Auth().Return(mockAuthService)
	mockAuthService.EXPECT().UserInfo(gomock.Any()).Return(&dto.AuthUserInfoResponse{}, nil)

	// 测试登陆的情况
	w = doRequest("GET", "/auth/user_info", nil)
	assert.Equal(t, http.StatusOK, w.Code)

}

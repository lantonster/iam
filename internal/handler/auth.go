package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/ginkit"
	"github.com/lantonster/iam/internal/dto"
	"github.com/lantonster/iam/internal/service"
)

type AuthHandler struct {
	service service.Service
}

func newAuthHandler(service service.Service) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

// Login godoc
//
//	@Summary	用户登陆
//	@Produce	json
//	@Param		query	query		dto.AuthLoginRequest	true	"用户登陆请求"
//	@Success	200		{object}	ginkit.SwaggerResponse{data=dto.AuthLoginResponse}
//	@Failure	400		{object}	ginkit.SwaggerResponseInvalidParam{}
//	@Router		/auth/login [get]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.AuthLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		ginkit.ResponseInvalidParam(c, err)
		return
	}

	res, err := h.service.Auth().Login(c, &req)
	ginkit.Response(c, res, err)
}

// UserInfo godoc
//
//	@Summary	获取当前登陆用户信息
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Success	200	{object}	ginkit.SwaggerResponse{data=dto.AuthUserInfoResponse}
//	@Failure	401	{object}	ginkit.SwaggerResponseUnauthorized{}
//	@Router		/auth/user_info [get]
func (h *AuthHandler) UserInfo(c *gin.Context) {
	res, err := h.service.Auth().UserInfo(c)
	ginkit.Response(c, res, err)
}

// UsernameAvailable godoc
//
//	@Summary	验证用户名可用性
//	@Produce	json
//	@Param		query	query		dto.AuthUsernameAvailableRequest	true	"验证用户名可用性"
//	@Success	200		{object}	ginkit.SwaggerResponse{}
//	@Failure	400		{object}	ginkit.SwaggerResponseInvalidParam{}
//	@Router		/auth/username_available [get]
func (h *AuthHandler) UsernameAvailable(c *gin.Context) {
	var req dto.AuthUsernameAvailableRequest
	if err := c.ShouldBind(&req); err != nil {
		ginkit.ResponseInvalidParam(c, err)
		return
	}

	err := h.service.Auth().UsernameAvailable(c, &req)
	ginkit.Response(c, nil, err)
}

// SendCode godoc
//
//	@Summary	发送验证码
//	@Produce	json
//	@Param		body	body		dto.AuthSendCodeRequest	true	"发送验证码"
//	@Success	200		{object}	ginkit.SwaggerResponse{}
//	@Failure	400		{object}	ginkit.SwaggerResponseInvalidParam{}
//	@Router		/auth/send_code [post]
func (h *AuthHandler) SendCode(c *gin.Context) {
	var req dto.AuthSendCodeRequest
	fmt.Printf("c.Request: %+v\n", c.Request)
	fmt.Printf("c.Request.Body: %+v\n", c.Request.Body)
	if err := c.ShouldBind(&req); err != nil {
		ginkit.ResponseInvalidParam(c, err)
		return
	}

	err := h.service.Auth().SendCode(c, &req)
	ginkit.Response(c, nil, err)
}

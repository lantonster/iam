package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lantonster/ginkit"
	"github.com/lantonster/iam/internal/dto"
	"github.com/lantonster/iam/internal/service"
)

type AuthHandler struct {
	service *service.Service
}

func newAuthHandler(service *service.Service) *AuthHandler {
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

	res, err := h.service.Auth.Login(c, &req)
	ginkit.Response(c, res, err)
}

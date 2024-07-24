package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lantonster/cerrors"
	"github.com/lantonster/ecodes"
	"github.com/lantonster/iam/internal/dto"
	"github.com/lantonster/iam/internal/repo"
	"github.com/lantonster/iam/pkg/utils"
)

type AuthService interface {
	// Login 登录
	Login(c *gin.Context, req *dto.AuthLoginRequest) (res *dto.AuthLoginResponse, err error)

	// UserInfo 获取用户信息
	UserInfo(c *gin.Context) (res *dto.AuthUserInfoResponse, err error)
}

type defaultAuthService struct {
	repo repo.Repo
}

func newDefaultAuthService(repo repo.Repo) *defaultAuthService {
	return &defaultAuthService{
		repo: repo,
	}
}

func (s *defaultAuthService) Login(c *gin.Context, req *dto.AuthLoginRequest) (res *dto.AuthLoginResponse, err error) {
	res = &dto.AuthLoginResponse{}

	// 通过用户名获取用户信息
	user, err := s.repo.User().GetUserByUsername(c, req.Username)
	if err != nil {
		return nil, err
	}

	// 验证密码
	if !user.ComparePassword(req.Password) {
		return nil, cerrors.WithCode(ecodes.IAM_PASSWORD_ERROR, "password error")
	}

	// 生成 token
	if res.Token, err = utils.GenerateToken(user.Id, user.Username); err != nil {
		return nil, cerrors.Wrap(err, "generate token")
	}

	return res, nil
}

func (s *defaultAuthService) UserInfo(c *gin.Context) (res *dto.AuthUserInfoResponse, err error) {
	res = &dto.AuthUserInfoResponse{
		UserId:   utils.GetUserIdFromContext(c),
		Username: utils.GetUsernameFromContext(c),
	}

	return res, nil
}

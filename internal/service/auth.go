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
	// 登录
	Login(c *gin.Context, req *dto.AuthLoginRequest) (res *dto.AuthLoginResponse, err error)

	// 获取用户信息
	UserInfo(c *gin.Context) (res *dto.AuthUserInfoResponse, err error)

	// 用户名是否可用
	UsernameAvailable(c *gin.Context, req *dto.AuthUsernameAvailableRequest) error
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

// UsernameAvailable 方法用于检查用户名是否可用
func (s *defaultAuthService) UsernameAvailable(c *gin.Context, req *dto.AuthUsernameAvailableRequest) error {
	// 检查用户名长度，如果小于 1 或大于 30，则返回错误
	if l := len(req.Username); l < 1 || 30 < l {
		return cerrors.WithCode(ecodes.IAM_INVALID_USERNAME_LENGTH, "the length of username should within the range of 1 to 30")
	}

	// 检查用户名格式，如果不符合要求（不是由字母、数字和下划线组成），则返回错误
	if !utils.ValidateUsernameFormat(req.Username) {
		return cerrors.WithCode(ecodes.IAM_INVALID_USERNAME_FORMAT, "the username can only be composed of letters, numbers and underscores")
	}

	// 检查用户名是否重复，如果重复，则返回错误
	if duplicated, err := s.repo.User().CheckUserNameDuplication(c, req.Username); err != nil {
		return err
	} else if duplicated {
		return cerrors.WithCode(ecodes.IAM_USERNAME_ALREADY_EXISTS, "the username already exists")
	}

	return nil
}

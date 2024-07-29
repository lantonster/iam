package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/cerrors"
	"github.com/lantonster/ecodes"
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/dto"
	"github.com/lantonster/iam/internal/model"
	"github.com/lantonster/iam/internal/repo"
	"github.com/lantonster/iam/pkg/utils"
	"gopkg.in/gomail.v2"
)

type AuthService interface {
	// 登录
	Login(c *gin.Context, req *dto.AuthLoginRequest) (res *dto.AuthLoginResponse, err error)

	// 获取用户信息
	UserInfo(c *gin.Context) (res *dto.AuthUserInfoResponse, err error)

	// 用户名是否可用
	UsernameAvailable(c *gin.Context, req *dto.AuthUsernameAvailableRequest) error

	// 发送验证码
	SendCode(c *gin.Context, req *dto.AuthSendCodeRequest) error
}

type defaultAuthService struct {
	email *config.EmailAuth
	repo  repo.Repo
}

func newDefaultAuthService(conf *config.Config, repo repo.Repo) *defaultAuthService {
	return &defaultAuthService{
		email: conf.Email,
		repo:  repo,
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

func (s *defaultAuthService) UsernameAvailable(c *gin.Context, req *dto.AuthUsernameAvailableRequest) error {
	// 检查用户名是否合法
	if err := model.IsUsernameValid(req.Username); err != nil {
		return err
	}

	// 检查用户名是否重复，如果重复，则返回错误
	if duplicated, err := s.repo.User().CheckUserNameDuplication(c, req.Username); err != nil {
		return err
	} else if duplicated {
		return cerrors.WithCode(ecodes.IAM_USERNAME_ALREADY_EXISTS, "the username already exists")
	}

	return nil
}

func (s *defaultAuthService) SendCode(c *gin.Context, req *dto.AuthSendCodeRequest) error {
	// 生成验证码
	code, err := s.repo.VerificationCode().GenerateCode(c, req.Email)
	if err != nil {
		return cerrors.Wrap(err, "generate verification code")
	}

	// 获取邮件相关配置
	vc := s.email.VerificationCode
	// 创建邮件消息
	m := gomail.NewMessage()
	// 设置发件人
	m.SetHeader("From", vc.Sender)
	// 设置收件人
	m.SetHeader("To", req.Email)
	// 设置邮件主题
	m.SetHeader("Subject", vc.Subject)

	// 构建邮件内容，将验证码插入到指定内容模板中
	content := fmt.Sprintf(vc.Content, code)
	m.SetBody("text/html", content)

	// 创建发送器并尝试发送邮件
	d := gomail.NewDialer(s.email.SmtpHost, s.email.SmtpPort, s.email.SmtpUsername, s.email.SmtpPassword)
	if err := d.DialAndSend(m); err != nil {
		return cerrors.WithCode(ecodes.IAM_SEND_VERIFICATION_CODE_FAILED, err.Error())
	}

	return nil
}

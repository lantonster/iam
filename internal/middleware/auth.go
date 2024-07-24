package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/cerrors"
	"github.com/lantonster/ecodes"
	"github.com/lantonster/ginkit"
	"github.com/lantonster/iam/pkg/utils"
)

var publicRoutes = []string{
	"/auth/login",
}

// LoginRequiredMiddleware 是一个 Gin 中间件函数，用于处理请求的登录验证和授权
func LoginRequiredMiddleware(c *gin.Context) {
	// 如果当前请求的路径是公开路由，则直接放行，执行后续的处理函数
	for _, rout := range publicRoutes {
		if c.Request.URL.Path == rout {
			c.Next()
			return
		}
	}

	// 从请求头中获取 'Authorization' 的值
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ginkit.ResponsesUnauthorized(c, cerrors.WithCode(ecodes.IAM_AUTHENTICATION_TOKEN_NOT_PROVIDED, "authentication token was not provided"))
		return
	}

	// 按照空格分割 'Authorization' 值
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ginkit.ResponsesUnauthorized(c, cerrors.WithCode(ecodes.IAM_AUTHENTICATION_TOKEN_INVALID_FORMAT, "authentication token format is invalid"))
		return
	}

	// 解析令牌字符串
	tokenString := parts[1]
	token, err := utils.ParseToken(tokenString)
	if err != nil {
		ginkit.ResponsesUnauthorized(c, cerrors.WithCode(ecodes.IAM_AUTHENTICATION_TOKEN_INVALID, err.Error()))
		return
	}

	// 将解析出的令牌中的用户 id、用户名设置到当前请求的上下文中
	utils.SetUserIdToContext(c, token.UserId)
	utils.SetUsernameToContext(c, token.Username)

	c.Next()
}

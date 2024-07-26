package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lantonster/iam/internal/handler"
	"github.com/lantonster/iam/internal/middleware"
)

func initAuthRouter(r *gin.Engine, h *handler.AuthHandler) {
	g := r.Group("/auth")
	{
		// 登录
		g.GET("/login", h.Login)

		// 获取当前登陆用户信息（需要登陆校验）
		g.GET("/user_info", middleware.LoginRequiredMiddleware, h.UserInfo)

		// 验证用户名可用性
		g.GET("/username_available", h.UsernameAvailable)

		// 验证邮箱可用性
		g.GET("/email_available")

		// 密码强度检查
		g.POST("/check_password")

		// 发送验证码
		g.POST("/send_code")

		// 注册
		g.POST("/register")
	}
}

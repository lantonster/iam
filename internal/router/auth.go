package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lantonster/iam/internal/handler"
)

func initAuthRouter(r *gin.Engine, h *handler.AuthHandler) {
	g := r.Group("/auth")
	{
		// 登录
		g.GET("/login", h.Login)

		g.GET("/logout")

		// 获取当前登陆用户信息
		g.GET("/user_info", h.UserInfo)

		g.GET("/register")
	}
}

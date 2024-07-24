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

		g.GET("/user_info")

		g.GET("/register")
	}
}

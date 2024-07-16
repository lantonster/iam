package handler

import "github.com/gin-gonic/gin"

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// Login godoc
//
//	@Summary	用户登陆
//	@Produce	json
//	@Param		username	query	string	true	"用户名" maxlength(20)
//	@Param		password	query	string	true	"密码" maxlength(60)
//	@Router		/auth/login [get]
func (h *AuthHandler) Login(c *gin.Context) {

}

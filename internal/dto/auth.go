package dto

type AuthLoginRequest struct {
	Username string `form:"username" binding:"required" default:"admin"`  // 用户名
	Password string `form:"password" binding:"required" default:"123456"` // 密码
}

type AuthLoginResponse struct {
	Token string `json:"token" default:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE4NzI3ODUsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.ZNgtQlyfVacyBg_ZouF4C3CpiMVxIaWXrh_a1i-OiAw"` // token
}

type AuthUserInfoResponse struct {
	UserId   int64  `json:"userId" default:"1"`       // 用户 id
	Username string `json:"username" default:"admin"` // 用户名
}

type AuthUsernameAvailableRequest struct {
	Username string `form:"username" binding:"required" default:"admin"`
}

type AuthSendCodeRequest struct {
	Email string `json:"email" binding:"required" default:"example@example.com"`
}

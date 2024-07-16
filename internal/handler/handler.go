package handler

type Handler struct {
	AuthHandler *AuthHandler
}

func NewHandler(authHandler *AuthHandler) *Handler {
	return &Handler{
		AuthHandler: authHandler,
	}
}

package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

const (
	ContextKeyUserId   = "user_id"
	ContextKeyUsername = "username"
)

func SetUserIdToContext(c *gin.Context, userId int64) {
	c.Set(ContextKeyUserId, userId)
}

func GetUserIdFromContext(c *gin.Context) int64 {
	userId, _ := c.Get(ContextKeyUserId)
	return cast.ToInt64(userId)
}

func SetUsernameToContext(c *gin.Context, username string) {
	c.Set(ContextKeyUsername, username)
}

func GetUsernameFromContext(c *gin.Context) string {
	username, _ := c.Get(ContextKeyUsername)
	return cast.ToString(username)
}

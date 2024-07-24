package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lantonster/cerrors"
	"github.com/spf13/cast"
)

type TokenBody struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}

const (
	// TokenExpireDuration 令牌过期时间
	TokenExpireDuration = time.Hour * 24
	// TokenSecret 令牌密钥
	TokenSecret = "iam_token_secret"
)

const (
	// TokenKeyUserId 用户 id
	TokenKeyUserId = "user_id"
	// TokenKeyUsername 用户名
	TokenKeyUsername = "username"
	// TokenKeyExpireTime 过期时间
	TokenKeyExpireTime = "exp"
)

// GenerateToken 生成令牌
func GenerateToken(userId int64, username string) (string, error) {
	// 创建一个新的令牌对象
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置令牌的声明（Claims）
	claims := token.Claims.(jwt.MapClaims)
	claims[TokenKeyUserId] = userId
	claims[TokenKeyUsername] = username
	claims[TokenKeyExpireTime] = time.Now().Add(TokenExpireDuration).Unix()

	// 使用一个密钥来签名令牌
	tokenString, err := token.SignedString([]byte(TokenSecret))
	if err != nil {
		return "", cerrors.Wrap(err, "generate token")
	}
	return tokenString, nil
}

// ParseToken 解析令牌
func ParseToken(tokenString string) (*TokenBody, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &TokenBody{
			UserId:   cast.ToInt64(claims[TokenKeyUserId]),
			Username: cast.ToString(claims[TokenKeyUsername]),
		}, nil
	}

	return nil, cerrors.New("invalid token")
}

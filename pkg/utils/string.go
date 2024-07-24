package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword 使用 bcrypt 算法对密码进行加密。
// 它接受一个密码和一个盐值作为输入，返回加密后的密码字符串和可能的错误。
// 加盐是为了增加密码的复杂度，使加密更加安全。
//
// 参数:
//   - password: 需要加密的明文密码。
//   - salt: 用于增强加密安全性的盐值。
//
// 返回值:
//   - 加密后的密码字符串。
//   - 可能发生的错误。
func HashPassword(password, salt string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// VerifyPassword 用于验证用户输入的密码是否与存储的哈希密码匹配。
// 它首先将密码和盐拼接在一起，然后使用bcrypt库比较拼接后的密码和存储的哈希密码。
//
// 参数:
//   - password: 用户输入的明文密码。
//   - salt: 用于加强密码安全性的盐值。
//   - hashedPassword: 存储的密码哈希值。
//
// 返回值:
//   - 如果输入的密码和盐经过哈希处理后与存储的哈希密码匹配，则返回true；否则返回false。
func VerifyPassword(password, salt, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt)) == nil
}

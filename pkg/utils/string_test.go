package utils

import (
	"testing"
)

func TestValidateUsernameFormat(t *testing.T) {
	t.Run("测试有效的用户名", func(t *testing.T) {
		validUsernames := []string{"a", "valid_user", "user_123", "user_name", ""}
		for _, username := range validUsernames {
			if !ValidateUsernameFormat(username) {
				t.Errorf("Expected ValidateUserName(%q) to be true, but got false", username)
			}
		}
	})

	t.Run("测试无效的用户名", func(t *testing.T) {
		invalidUsernames := []string{"invalid@user", "user!name", "$", "/t", "[]", "中文", "!", "@"}
		for _, username := range invalidUsernames {
			if ValidateUsernameFormat(username) {
				t.Errorf("Expected ValidateUserName(%q) to be false, but got true", username)
			}
		}
	})
}

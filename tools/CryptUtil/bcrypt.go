package CryptUtil

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateBcryptPassword 生成密码
func GenerateBcryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CompareHashAndPassword 验证新旧密码
func CompareHashAndPassword(hashPassword string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) string {
	salt := 10

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), salt)

	return string(hash)
}

func ComparePass(hashPassword, password []byte) bool {
	hash, pass := []byte(hashPassword), []byte(password)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}

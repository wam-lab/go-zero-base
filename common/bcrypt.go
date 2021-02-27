package common

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(source string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(password), err
}

func CheckPassword(crypt, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(crypt), []byte(plain))
	return err == nil
}

package models

import "golang.org/x/crypto/bcrypt"

type LoginUser struct {
	Name     string
	Password []byte
}

func (u *LoginUser) GetPassHash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (u *LoginUser) VerifyPass(PassHash string) bool {
	bytePassHash := []byte(PassHash)

	err := bcrypt.CompareHashAndPassword(bytePassHash, u.Password)
	if err != nil {
		return false
	}

	return true
}

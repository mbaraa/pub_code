package services

import (
	"errors"
	"spendings/session"
)

type AuthService struct{}

func (a *AuthService) CheckPassword(pw string) error {
	if session.CheckPassword(pw) {
		return nil
	}
	return errors.New("password doesn't match")
}

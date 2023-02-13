package model

import (
	errors2 "github.com/Selepok/calendar/internal/errors"
)

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (auth *Auth) OK() error {
	if len(auth.Login) == 0 {
		return errors2.ErrMissingField("login")
	}
	if len(auth.Password) == 0 {
		return errors2.ErrMissingField("password")
	}

	return nil
}

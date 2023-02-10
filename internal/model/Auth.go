package model

import "errors"

type Auth struct {
	Login    string
	Password string
}

func (auth *Auth) OK() error {
	if len(auth.Login) < 6 {
		return errors.New("the length of the Login should be at least 6 symbols")
	}
	if len(auth.Password) < 6 {
		return errors.New("the length of the Password should be at least 6 symbols")
	}

	return nil
}

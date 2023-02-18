package model

import (
	errors2 "github.com/Selepok/calendar/internal/errors"
	"time"
)

type CreateUser struct {
	TimeZone string `json:"timezone"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *CreateUser) OK() (err error) {
	if len(u.Login) == 0 {
		return errors2.ErrMissingField("login")
	}
	if len(u.Password) == 0 {
		return errors2.ErrMissingField("password")
	}

	if len(u.TimeZone) == 0 {
		return errors2.ErrMissingField("timezone")
	}
	_, err = time.LoadLocation(u.TimeZone)
	if err != nil {
		return errors2.TimezoneNotValid(u.TimeZone)
	}

	return nil
}

type User struct {
	TimeZone string `json:"timezone"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *User) OK() (err error) {
	if len(u.Login) == 0 {
		return errors2.ErrMissingField("login")
	}

	if len(u.TimeZone) == 0 {
		return errors2.ErrMissingField("timezone")
	}
	_, err = time.LoadLocation(u.TimeZone)
	if err != nil {
		return errors2.TimezoneNotValid(u.TimeZone)
	}

	return nil
}

package errors

import (
	"fmt"
)

type ErrMissingField string

func (e ErrMissingField) Error() string {
	return string(e) + " is required"
}

type NoUserFound string

func (e NoUserFound) Error() string {
	return fmt.Sprintf("There is no user with login: %s.", string(e))
}

type UserCreationIssue struct {
}

func (e UserCreationIssue) Error() string {
	return "Something went wrong while creating user."
}

type IncorrectPassword string

func (e IncorrectPassword) Error() string {
	return fmt.Sprintf("The password for user %s is incorrect.", string(e))
}

type GenerateTokenIssue struct {
}

func (e GenerateTokenIssue) Error() string {
	return "the error occurred while generating token"
}

package constants

import "errors"

var (
	// global
	ErrUnexpected = errors.New("unexpected error")

	// entity
	ErrUserNotFound = errors.New("user not found")

	// config
	ErrLoadConfig  = errors.New("failed to load config file")
	ErrParseConfig = errors.New("failed to parse env to config struct")
	ErrEmptyVar    = errors.New("required variabel environment is empty")
	ErrTokenExpired = errors.New("token expired")
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenDoesNotExist = errors.New("token does not exist")
	ErrDatabaseUpdate = errors.New("failed to update database")
	ErrInvalidResetToken = errors.New("invalid reset token")
)

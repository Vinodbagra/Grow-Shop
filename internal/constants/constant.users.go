package constants

import "errors"

const (
	CtxAuthenticatedUserKey = "CtxAuthenticatedUserKey"
	AdminID                 = 1
	UserID                  = 2
)

var (
	ListGender             = []string{"male", "female"}
	ErrInvalidUUIDFormat = errors.New("invalid  uuid format")
)

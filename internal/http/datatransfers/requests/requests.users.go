package requests

import (
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
)

// General Request
type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,containsany=!@#$%^&*()?"`
}

// Mapping General Request to Domain User
func (user UserRequest) ToV1Domain() *V1Domains.UserDomain {
	return &V1Domains.UserDomain{
		Email:    user.Email,
		Password: user.Password,
	}
}

// Send OTP Request
type UserForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// Verify OTP Code
type UserResetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,containsany=!@#$%^&*()?"`
	ResetToken  string `json:"reset_token" validate:"required"`
}

func (user UserResetPasswordRequest) ToV1Domain() *V1Domains.UserDomain {
	return &V1Domains.UserDomain{
		Email:    user.Email,
		Password: user.Password,
	}
}

// Login Request
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,containsany=!@#$%^&*()?"`
}

// Mapping Login Request to Domain User
func (u *UserLoginRequest) ToV1Domain() *V1Domains.UserDomain {
	return &V1Domains.UserDomain{
		Email:    u.Email,
		Password: u.Password,
	}
}

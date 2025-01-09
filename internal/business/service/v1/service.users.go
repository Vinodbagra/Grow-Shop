package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/pkg/helpers"
	"github.com/snykk/grow-shop/pkg/mailer"
)

type userservice struct {
	repo   V1Domains.UserRepository
	mailer mailer.OTPMailer
}

type Userservice interface {
	Store(ctx context.Context, inDom *V1Domains.UserDomain) (outDom V1Domains.UserDomain, statusCode int, err error)
	Login(ctx context.Context, inDom *V1Domains.UserDomain) (outDom V1Domains.UserDomain, statusCode int, err error)
	SendOTP(ctx context.Context, email string) (otpCode string, statusCode int, err error)
	VerifOTP(ctx context.Context, email string, userOTP string, otpRedis string) (statusCode int, err error)
	ActivateUser(ctx context.Context, email string) (statusCode int, err error)
	GetByEmail(ctx context.Context, email string) (outDom V1Domains.UserDomain, statusCode int, err error)
}

func NewUserservice(repo V1Domains.UserRepository, mailer mailer.OTPMailer) Userservice {
	return &userservice{
		repo:   repo,
		mailer: mailer,
	}
}

func (user *userservice) Store(ctx context.Context, inDom *V1Domains.UserDomain) (outDom V1Domains.UserDomain, statusCode int, err error) {
	inDom.Password, err = helpers.GenerateHash(inDom.Password)
	if err != nil {
		return V1Domains.UserDomain{}, http.StatusInternalServerError, err
	}

	inDom.CreatedAt = time.Now().In(constants.GMT7)
	fmt.Println(time.Now().In(constants.GMT7))
	err = user.repo.Store(ctx, inDom)
	if err != nil {
		return V1Domains.UserDomain{}, http.StatusInternalServerError, err
	}

	outDom, err = user.repo.GetByEmail(ctx, inDom)
	if err != nil {
		return V1Domains.UserDomain{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusCreated, nil
}

func (user *userservice) Login(ctx context.Context, inDom *V1Domains.UserDomain) (outDom V1Domains.UserDomain, statusCode int, err error) {
	userDomain, err := user.repo.GetByEmail(ctx, inDom)
	if err != nil {
		return V1Domains.UserDomain{}, http.StatusUnauthorized, errors.New("invalid email or password") // for security purpose better use generic error message
	}

	if !helpers.ValidateHash(inDom.Password, userDomain.Password) {
		return V1Domains.UserDomain{}, http.StatusUnauthorized, errors.New("invalid email or password")
	}

	// if userDomain.RoleID == constants.AdminID {
	// 	userDomain.Token, err = user.jwtService.GenerateToken(userDomain.ID, true, userDomain.Email, userDomain.Password)
	// } else {
	// 	userDomain.Token, err = user.jwtService.GenerateToken(userDomain.ID, false, userDomain.Email, userDomain.Password)
	// }

	// if err != nil {
	// 	return V1Domains.UserDomain{}, http.StatusInternalServerError, err
	// }

	return userDomain, http.StatusOK, nil
}

func (user *userservice) SendOTP(ctx context.Context, email string) (otpCode string, statusCode int, err error) {
	_, err = user.repo.GetByEmail(ctx, &V1Domains.UserDomain{Email: email})
	if err != nil {
		return "", http.StatusNotFound, errors.New("email not found")
	}

	code, err := helpers.GenerateOTPCode(6)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	if err = user.mailer.SendOTP(code, email); err != nil {
		return "", http.StatusInternalServerError, err
	}

	return code, http.StatusOK, nil
}

func (user *userservice) VerifOTP(ctx context.Context, email string, userOTP string, otpRedis string) (statusCode int, err error) {
	_, err = user.repo.GetByEmail(ctx, &V1Domains.UserDomain{Email: email})
	if err != nil {
		return http.StatusNotFound, errors.New("email not found")
	}

	if otpRedis != userOTP {
		return http.StatusBadRequest, errors.New("invalid otp code")
	}

	return http.StatusOK, nil
}

func (u *userservice) ActivateUser(ctx context.Context, email string) (statusCode int, err error) {
	user, err := u.repo.GetByEmail(ctx, &V1Domains.UserDomain{Email: email})
	if err != nil {
		return http.StatusNotFound, errors.New("email not found")
	}

	if err = u.repo.ChangeActiveUser(ctx, &V1Domains.UserDomain{UserID: user.UserID}); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (u *userservice) GetByEmail(ctx context.Context, email string) (outDom V1Domains.UserDomain, statusCode int, err error) {
	user, err := u.repo.GetByEmail(ctx, &V1Domains.UserDomain{Email: email})
	if err != nil {
		return V1Domains.UserDomain{}, http.StatusNotFound, errors.New("email not found")
	}

	return user, http.StatusOK, nil
}

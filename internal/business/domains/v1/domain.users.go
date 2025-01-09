package v1

import (
	"context"
	"time"
)

type UserDomain struct {
	UserID       string
	UserName     string
	Email        string
	MobileNo     string
	Address      string
	Password     string
	BusinessName string
	Gender       string
	Shops        []string
	LicenseID    string
	Images       []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepository interface {
	Store(ctx context.Context, inDom *UserDomain) (err error)
	GetByEmail(ctx context.Context, inDom *UserDomain) (outDomain UserDomain, err error)
	ChangeActiveUser(ctx context.Context, inDom *UserDomain) (err error)
}

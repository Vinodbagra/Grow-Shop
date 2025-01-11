package v1

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UserDomain struct {
	UserID       uuid.UUID
	UserName     string
	Email        string
	MobileNo     sql.NullString
	Address      sql.NullString
	Password     string
	BusinessName sql.NullString
	Gender       string
	Shops        []string
	LicenseID    uuid.UUID
	Images       []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepository interface {
	Store(ctx context.Context, inDom *UserDomain) (err error)
	GetByEmail(ctx context.Context, inDom *UserDomain) (outDomain UserDomain, err error)
	ChangeActiveUser(ctx context.Context, inDom *UserDomain) (err error)
	GetUserByID(ctx context.Context, inDom *UserDomain) (outDomain UserDomain, err error)
}

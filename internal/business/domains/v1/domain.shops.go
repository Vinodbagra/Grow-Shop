package v1

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// have to write nessesary things
type ShopDomain struct {
	ShopName     string
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

type ShopRepository interface {
	Store(ctx context.Context, inDom *UserDomain) (err error)
}

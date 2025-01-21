package v1

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// have to write nessesary things
type ShopDomain struct {
	ShopID          uuid.UUID
	UserID          uuid.UUID
	ShopName        string
	ShopAddress     string
	ShopImages      []string
	ShopDescription string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ShopRepository interface {
	Store(ctx context.Context, inDom *UserDomain) (err error)
}

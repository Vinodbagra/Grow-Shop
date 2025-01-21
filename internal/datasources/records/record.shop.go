package records

import (
	"time"

	"github.com/google/uuid"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
)

type Shops struct {
	ShopID          uuid.UUID `db:"shop_id"`          // Shop ID (UUID primary key)
	UserID          uuid.UUID `db:"user_id"`          // Foreign key referencing users table
	ShopName        string    `db:"shop_name"`        // Shop name
	ShopAddress     string    `db:"shop_address"`     // Shop address (optional)
	ShopImages      []string  `db:"shop_images"`      // List of shop image URLs
	ShopDescription string    `db:"shop_description"` // Shop description (optional)
	CreatedAt       time.Time `db:"created_at"`       // Record creation time
	UpdatedAt       time.Time `db:"updated_at"`       // Last update time
}

func FromShopsV1Domain(s *V1Domains.ShopDomain) Shops {
	return Shops{
		ShopID: s.ShopID,
		UserID: s.UserID,
		ShopName: s.ShopName,
		ShopAddress: s.ShopAddress,
		ShopImages: s.ShopImages,
		ShopDescription: s.ShopDescription,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func (s *Shops) ToV1Domain() V1Domains.ShopDomain {
	return V1Domains.ShopDomain{
		ShopName: s.ShopName,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

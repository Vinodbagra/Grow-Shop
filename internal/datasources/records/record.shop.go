package records

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
)

type Shops struct {
	ShopName     string         `db:"shop_name"`
	UserID       uuid.UUID      `db:"user_id"`       // UUID primary key
	UserName     string         `db:"user_name"`     // User name
	MobileNo     sql.NullString `db:"mobile_no"`     // Mobile number
	Address      sql.NullString `db:"address"`       // Address (optional)
	UserPassword string         `db:"password"`      // Encrypted password
	BusinessName sql.NullString `db:"business_name"` // Business name (optional)
	Gender       string         `db:"gender"`        // Gender (default 'Other')
	Shops        pq.StringArray `db:"shops"`         // List of shop UUIDs
	LicenseID    uuid.UUID      `db:"license_id"`    // Single UUID for license
	Images       pq.StringArray `db:"images"`        // List of image URLs
	CreatedAt    time.Time      `db:"created_at"`    // Record creation time
	UpdatedAt    time.Time      `db:"updated_at"`    // Last update time
}

func (s *Shops) ToV1Domain() V1Domains.ShopDomain {
	return V1Domains.ShopDomain{
		ShopName: s.ShopName,
		UserName:  s.UserName,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func FromShopV1Domain(s *V1Domains.ShopDomain) Shops {
	return Shops{
		UserName:     s.UserName,
		ShopName:     s.ShopName,
		MobileNo:     s.MobileNo,
		Address:      s.Address,
		BusinessName: s.BusinessName,
		Gender:       s.Gender,
		Shops:        s.Shops,
		LicenseID:    s.LicenseID,
		Images:       s.Images,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}

package records

import (
	"time"

	"github.com/google/uuid"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
)

type License struct {
	LicenseID   uuid.UUID `db:"license_id"`
	UserID      uuid.UUID `db:"user_id"`
	LicenseType string    `db:"license_type"`
	ShopLimit   int       `db:"shop_limit"`
	Validity    time.Time `db:"validity"`
	CreatedAt   time.Time `db:"created_at"` // Record creation time
	UpdatedAt   time.Time `db:"updated_at"`
}

func FromLicenseV1Domain(u *V1Domains.LicenseDomain) License {
	return License{
		LicenseID:   u.LicenseID,
		UserID:       u.UserID,
		LicenseType:  u.LicenseType,
		ShopLimit:    u.ShopLimit,
		Validity:     u.Validity,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

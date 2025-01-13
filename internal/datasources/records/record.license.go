package records

import (
	"time"

	"github.com/google/uuid"
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

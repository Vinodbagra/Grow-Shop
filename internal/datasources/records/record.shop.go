package records

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
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

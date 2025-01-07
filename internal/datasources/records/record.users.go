package records

import (
	"time"
)

type Users struct {
	UserID       string    `db:"user_id"`          // UUID primary key
	UserName     string    `db:"user_name"`        // User name
	Email        string    `db:"email"`           // Email (unique)
	MobileNo     string    `db:"mobile_no"`       // Mobile number
	Address      string    `db:"address"`         // Address (optional)
	UserPassword string    `db:"password"`        // Encrypted password
	BusinessName string    `db:"business_name"`   // Business name (optional)
	Gender       string    `db:"gender"`          // Gender (default 'Other')
	Shops        []string  `db:"shops"`           // List of shop UUIDs
	LicenseID    string    `db:"license_id"`      // Single UUID for license
	Images       []string  `db:"images"`          // List of image URLs
	CreatedAt    time.Time `db:"created_at"`      // Record creation time
	UpdatedAt    time.Time `db:"updated_at"`      // Last update time
}


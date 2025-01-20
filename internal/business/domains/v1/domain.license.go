package v1

import (
	"time"

	"github.com/google/uuid"
)

type LicenseDomain struct {
	LicenseID    uuid.UUID
	UserID       uuid.UUID
	DurationType string
	Duration     int
	LicenseType  string
	ShopLimit    int
	Validity     time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

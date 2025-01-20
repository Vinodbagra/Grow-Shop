package v1

import (
	"time"

	"github.com/google/uuid"
)

type TokenDomain struct {
	Token     uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

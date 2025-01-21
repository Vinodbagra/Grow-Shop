package records

import (
	"time"

	"github.com/google/uuid"
)

type Tokens struct {
	Token     uuid.UUID    `db:"token"`
	UserID    uuid.UUID    `db:"user_id"`
	CreatedAt time.Time    `db:"created_at"` // Record creation time
	UpdatedAt time.Time    `db:"updated_at"`
}



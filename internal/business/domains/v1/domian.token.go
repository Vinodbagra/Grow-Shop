package v1

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type TokenDomain struct {
	Token     string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TokenRepository interface {
	CreateToken(ctx context.Context, userID uuid.UUID) (token uuid.UUID,err error)
	ValidateToken(ctx context.Context, token uuid.UUID) (flag bool, err error)
}

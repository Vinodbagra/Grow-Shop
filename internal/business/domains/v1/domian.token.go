package v1

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type TokenDomain struct {
	Token     uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TokenRepository interface {
	CreateToken(ctx context.Context, userID uuid.UUID) (token uuid.UUID,err error)
	ValidateToken(ctx context.Context, token uuid.UUID) (tokens TokenDomain,flag bool, err error)
}

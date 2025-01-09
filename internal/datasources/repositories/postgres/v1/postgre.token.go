package v1

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
	"github.com/snykk/grow-shop/internal/datasources/records"
)

type postgreTokenRepository struct {
	conn *sqlx.DB
}

func NewTokenRepository(conn *sqlx.DB) V1Domains.TokenRepository {
	return &postgreTokenRepository{
		conn: conn,
	}
}

func (r *postgreTokenRepository) CreateToken(ctx context.Context, userID uuid.UUID) (token uuid.UUID,err error) {
	tokens := records.Tokens{
		UserID:    userID,
	}
	row, err := r.conn.NamedQueryContext(ctx, `INSERT INTO tokens(token, user_id, created_at) VALUES (uuid_generate_v4(), :user_id, :created_at)`, tokens)
	err = row.Scan(&token)
	if err != nil {
		return uuid.Nil, err
	}

	return token, nil
}

func (r *postgreTokenRepository) ValidateToken(ctx context.Context, token uuid.UUID) (flag bool, err error) {
	tokens := records.Tokens{}
	err = r.conn.GetContext(ctx, &tokens, `SELECT * FROM users WHERE "email" = $1`, token)
	if err != nil {
		return false, err
	}
	if tokens.Token == token {
		return true, nil
	}

	return false,nil
}




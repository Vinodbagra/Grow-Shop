package v1

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/internal/datasources/records"
	"github.com/snykk/grow-shop/pkg/logger"
)

type postgreTokenRepository struct {
	conn *sqlx.DB
}

func NewTokenRepository(conn *sqlx.DB) V1Domains.TokenRepository {
	return &postgreTokenRepository{
		conn: conn,
	}
}

func (r *postgreTokenRepository) CreateToken(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	methodName := "postgreTokenRepository.CreateToken"
	logger.InfoF("function name %s recieved the request to create token", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)
    var existingToken records.Tokens

    // Check if a token already exists for the user
    err := r.conn.GetContext(ctx, &existingToken, `SELECT * FROM tokens WHERE user_id = $1`, userID)
    if err != nil {
        if err == sql.ErrNoRows {
            // No existing token, create a new one
            newToken := uuid.New()
            _, err = r.conn.ExecContext(ctx, 
                `INSERT INTO tokens (token, user_id, created_at) VALUES ($1, $2, NOW())`, 
                newToken, userID,
            )
            if err != nil {
                logger.ErrorF("error when inserting new token: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase}, err)
                return uuid.Nil, err
            }
            return newToken, nil
        }
        // Unexpected error
        logger.ErrorF("error when querying token: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase}, err)
        return uuid.Nil, err
    }

    // If token exists, check for expiration
    expirationDate := existingToken.CreatedAt.Add(30 * 24 * time.Hour) // 30 days from created_at
    if time.Now().After(expirationDate) {
        // Token expired, update with a new one
        updatedToken := uuid.New()
        _, err = r.conn.ExecContext(ctx, 
            `UPDATE tokens SET token = $1, created_at = NOW() WHERE user_id = $2`, 
            updatedToken, userID,
        )
        if err != nil {
            logger.ErrorF("error when updating token: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase}, err)
            return uuid.Nil, err
        }
        return updatedToken, nil
    }

    // Token exists and is valid
	logger.InfoF("function name %s successfully created token", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)
    return existingToken.Token, nil
}


func (r *postgreTokenRepository) ValidateToken(ctx context.Context, token uuid.UUID) (V1Domains.TokenDomain, bool, error) {
	tokens := records.Tokens{}
	tokenDomain := V1Domains.TokenDomain{}
	err := r.conn.GetContext(ctx, &tokens, `SELECT * FROM tokens WHERE "token" = $1`, token)
	if err != nil {
		if err == sql.ErrNoRows {
			// Token does not exist
			logger.ErrorF("Token not found: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
			return tokenDomain,false, constants.ErrTokenDoesNotExist
		}
		logger.ErrorF("error when executing query: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase}, err)
		return tokenDomain, false, err
	}

	// Check if token is expired
	expirationDate := tokens.CreatedAt.Add(30 * 24 * time.Hour) // 30 days from created_at
	if time.Now().After(expirationDate) {
		logger.InfoF("Token expired for user ID %v", logrus.Fields{"user_id": tokens.UserID})
		return tokenDomain, false, constants.ErrTokenExpired
	}

	// Check if the token matches
	tokenDomain.Token = tokens.Token
	tokenDomain.UserID = tokens.UserID
	tokenDomain.CreatedAt = tokens.CreatedAt
	tokenDomain.UpdatedAt = tokens.UpdatedAt
	if tokens.Token == token {
		return tokenDomain, true, nil
	}

	return tokenDomain, false, constants.ErrInvalidToken
}

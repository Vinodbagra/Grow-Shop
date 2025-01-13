package middlewares

import (
	"strings"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/snykk/grow-shop/internal/constants"
	V1Repository "github.com/snykk/grow-shop/internal/datasources/repositories/postgres/v1"
	V1Handler "github.com/snykk/grow-shop/internal/http/handlers/v1"
)

type AuthMiddleware struct {
	repo V1Repository.TokenRepository
}

func NewAuthMiddleware(conn *sqlx.DB, isAdmin bool) gin.HandlerFunc {
	return (&AuthMiddleware{
		repo: V1Repository.NewTokenRepository(conn),
	}).Handle
}

func (m *AuthMiddleware) Handle(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		V1Handler.NewAbortResponse(ctx, "missing authorization header")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		V1Handler.NewAbortResponse(ctx, "invalid header format")
		return
	}

	if headerParts[0] != "Bearer" {
		V1Handler.NewAbortResponse(ctx, "token must content bearer")
		return
	}

	token := headerParts[1]
	uuidToken, err := uuid.Parse(token)
	if err != nil {
		V1Handler.NewAbortResponse(ctx, "token is in invalid format")
		return
	}
	tokens, flag, err := m.repo.ValidateToken(ctx, uuidToken)
	if err == constants.ErrTokenDoesNotExist {
		V1Handler.NewAbortResponse(ctx, constants.ErrTokenDoesNotExist.Error())
		return
	} else if err == constants.ErrTokenExpired {
		V1Handler.NewAbortResponse(ctx, constants.ErrTokenExpired.Error())
		return
	}
	if err == constants.ErrInvalidToken {
		V1Handler.NewAbortResponse(ctx, constants.ErrInvalidToken.Error())
		return
	}

	if !flag {
		V1Handler.NewAbortResponse(ctx, "token is invalid")
		return
	}
	ctx.Set("userID", tokens.UserID)
	ctx.Next()
}

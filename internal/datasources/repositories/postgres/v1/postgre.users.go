package v1

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/internal/datasources/records"
	"github.com/snykk/grow-shop/pkg/logger"
)

type postgreUserRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) V1Domains.UserRepository {
	return &postgreUserRepository{
		conn: conn,
	}
}

func (r *postgreUserRepository) Store(ctx context.Context, inDom *V1Domains.UserDomain) (err error) {
	userRecord := records.FromUsersV1Domain(inDom)
	_, err = r.conn.NamedQueryContext(ctx, `INSERT INTO users(user_id,user_name, email, password, created_at) VALUES (uuid_generate_v4(), :user_name, :email, :password, :created_at)`, userRecord)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgreUserRepository) GetByEmail(ctx context.Context, inDom *V1Domains.UserDomain) (outDomain V1Domains.UserDomain, err error) {
	methodName := "postgreUserRepository.GetByEmail"
	logger.InfoF("function name %s recieved the request to get user by email", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)
	userRecord := records.FromUsersV1Domain(inDom)

	err = r.conn.GetContext(ctx, &userRecord, `SELECT * FROM users WHERE "email" = $1`, userRecord.Email)
	if err != nil {
		return V1Domains.UserDomain{}, err
	}
	logger.InfoF("function name %s successfully got user by email", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)

	return userRecord.ToV1Domain(), nil
}

func (r *postgreUserRepository) UpdatePassword(ctx context.Context, inDom *V1Domains.UserDomain) (err error) {
	userRecord := records.FromUsersV1Domain(inDom)

	_, err = r.conn.NamedQueryContext(ctx, `UPDATE users SET password = :password WHERE email = :email`, userRecord)
	if err != nil {
		logger.ErrorF("error when updating user password: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, err)
		return constants.ErrDatabaseUpdate
	}

	return nil
}

func (r *postgreUserRepository) GetUserByID(ctx context.Context, inDom *V1Domains.UserDomain) (outDomain V1Domains.UserDomain, err error) {
	methodName := "postgreUserRepository.GetByEmail"
	logger.InfoF("function name %s recieved the request to get user by id", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)
	userRecord := records.FromUsersV1Domain(inDom)

	err = r.conn.GetContext(ctx, &userRecord, `SELECT * FROM users WHERE "user_id" = $1`, userRecord.UserID)
	if err != nil {
		return V1Domains.UserDomain{}, err
	}
	logger.InfoF("function name %s successfully got user by id", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)

	return userRecord.ToV1Domain(), nil
}


// TODO
// create a function to update user data

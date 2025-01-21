package v1

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/internal/datasources/records"
	"github.com/snykk/grow-shop/pkg/logger"
)

type postgreShopRepository struct {
	conn *sqlx.DB
}

type ShopRepository interface {
	CreateShop(ctx context.Context, inDom *V1Domains.ShopDomain) (shopID uuid.UUID, err error)
}

func NewShopRepository(conn *sqlx.DB) ShopRepository {
	return &postgreShopRepository{
		conn: conn,
	}
}

func (r *postgreShopRepository) CreateShop(ctx context.Context, inDom *V1Domains.ShopDomain) (shopID uuid.UUID, err error) {
	shopRecord := records.FromShopsV1Domain(inDom)
	_, err = r.conn.NamedQueryContext(ctx, `INSERT INTO users(shop_name,user_id,user_name, created_at) VALUES (uuid_generate_v4(), :shop_name,user_name, :created_at)`, shopRecord)
	if err != nil {
		return V1Domains.ShopDomain{}.UserID, err
	}

	return shopRecord.ToV1Domain().UserID, nil // should return something else?
}

func (r *postgreShopRepository) GetShopByID(ctx context.Context, inDom *V1Domains.ShopDomain) (outDomain V1Domains.ShopDomain, err error) {
	methodName := "postgreShopRepository.GetByID"
	logger.InfoF("function name %s recieved the request to get user by id", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)
	shopRecord := records.FromShopV1Domain(inDom)

	err = r.conn.GetContext(ctx, &shopRecord, `SELECT * FROM users WHERE "user_id" = $1`, shopRecord.UserID)
	if err != nil {
		return V1Domains.ShopDomain{}, err
	}
	logger.InfoF("function name %s successfully got user by id", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)

	return shopRecord.ToV1Domain(), nil
}

// TODO
// create a function to update user data
func (r *postgreShopRepository) UpdateUserData(ctx context.Context, shopData *V1Domains.ShopDomain) (err error) {
	methodName := "postgreShopRepository.UpdateShopData"
	logger.InfoF("function name %s recieved the request to update user data", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)
	shopRecord := records.FromShopV1Domain(shopData)

	_, err = r.conn.NamedQueryContext(ctx, `UPDATE users SET mobile_no = :mobile_no, address = :address, business_name = :business_name, gender = :gender WHERE user_id = :user_id`, shopRecord)
	if err != nil {
		logger.ErrorF("error when updating shop details : %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, err)
		return constants.ErrDatabaseUpdate
	}
	logger.InfoF("function name %s successfully updated shop data", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)
	return nil
}

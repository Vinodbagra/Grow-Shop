package v1

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
	"github.com/snykk/grow-shop/internal/datasources/records"
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
	shopRecord := records.FromShopV1Domain(inDom)
	_, err = r.conn.NamedQueryContext(ctx, `INSERT INTO users(shop_name,user_id,user_name, created_at) VALUES (uuid_generate_v4(), :shop_name,user_name, :created_at)`, shopRecord)
	if err != nil {
		return err
	}

	return nil
}

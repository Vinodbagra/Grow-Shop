package v1

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/internal/datasources/records"
	"github.com/snykk/grow-shop/pkg/logger"
)

type postgreLicenseRepository struct {
	conn *sqlx.DB
}

func NewLicenseRepository(conn *sqlx.DB) LicenseRepository {
	return &postgreLicenseRepository{
		conn: conn,
	}
}

type LicenseRepository interface {
	CreateFreeLicense(ctx context.Context, userID uuid.UUID) (licenseID uuid.UUID, err error)
}

func (r *postgreLicenseRepository) CreateFreeLicense(ctx context.Context, userID uuid.UUID) (licenseID uuid.UUID, err error) {
	methodName := "postgreLicenseRepository.CreateFreeLicense"
	logger.InfoF("function name %s recieved the request to create free license", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)

	licenseData := records.License{
		LicenseID:   uuid.New(),
		UserID:      userID,
		ShopLimit:   1,
		Validity:    time.Now().Add(30 * 24 * time.Hour), // 30 days from now
	}
	_, err = r.conn.NamedQueryContext(ctx, `INSERT INTO license(license_id, user_id, validity, license_type, shop_limit) VALUES (license_id, :user_id, :validity, :license_type, :shop_limit, :created_at)`, licenseData)
	if err != nil {
		return uuid.Nil, err
	}
	return licenseData.LicenseID, nil
}

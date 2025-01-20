package v1

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/internal/datasources/records"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
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
	UpdateLicenseData(ctx context.Context, licenseData *V1Domains.LicenseDomain) (err error)
}

func (r *postgreLicenseRepository) CreateFreeLicense(ctx context.Context, userID uuid.UUID) (licenseID uuid.UUID, err error) {
	methodName := "postgreLicenseRepository.CreateFreeLicense"
	logger.InfoF("function name %s recieved the request to create free license", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)

	licenseData := records.License{
		LicenseID: uuid.New(),
		UserID:    userID,
		ShopLimit: 1,
		Validity:  time.Now().Add(30 * 24 * time.Hour), // 30 days from now
	}
	_, err = r.conn.NamedQueryContext(ctx, `INSERT INTO license(license_id, user_id, validity, license_type, shop_limit) VALUES (license_id, :user_id, :validity, :license_type, :shop_limit, :created_at)`, licenseData)
	if err != nil {
		return uuid.Nil, err
	}
	return licenseData.LicenseID, nil
}

func (r *postgreLicenseRepository) UpdateLicenseData(ctx context.Context, licenseData *V1Domains.LicenseDomain) (err error) {
	methodName := "postgreUserRepository.UpdateLicenseData"
	logger.InfoF("function name %s recieved the request to update license data", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)
	licenseRecord := records.FromLicenseV1Domain(licenseData)

	_, err = r.conn.NamedQueryContext(ctx, `UPDATE license SET license_type = :license_type, shop_limit = :shop_limit, validity = :validity WHERE license_id = :license_id`, licenseRecord)
	if err != nil {
		logger.ErrorF("error when updating user password: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, err)
		return constants.ErrDatabaseUpdate
	}
	logger.InfoF("function name %s successfully updated license data", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)
	return nil
}

package v1

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
	"github.com/snykk/grow-shop/internal/constants"
	V1Repository "github.com/snykk/grow-shop/internal/datasources/repositories/postgres/v1"
	"github.com/snykk/grow-shop/pkg/logger"
)

type licenseservice struct {
	licenseRepo V1Repository.LicenseRepository
}

type LicenseService interface {
	UpdateLicense(ctx context.Context, licenseData V1Domains.LicenseDomain) (statusCode int, err error)
}

func NewLicenseservice(licenseRepo V1Repository.LicenseRepository) LicenseService {
	return &licenseservice{
		licenseRepo: licenseRepo,
	}
}

func (license *licenseservice) UpdateLicense(ctx context.Context, licenseData V1Domains.LicenseDomain) (statusCode int, err error) {
	methodName := "licenseService.UpdateLicense"
	logger.InfoF("function name %s recieved the request to update license", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)

	// licenseID, err := license.licenseRepo.UpdateLicense(ctx, licenseData)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	logger.InfoF("function name %s successfully updated license", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)
	return http.StatusOK, nil
}

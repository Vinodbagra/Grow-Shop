package requests

import (
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
)

type UpdateLicenseRequest struct {
	LicenseType  string `json:"license_type" validate:"required,oneof=FREE SILVER GOLD other"`
	ShopLimit    int    `json:"shop_limit" validate:"required,min=1"`
	DurationType string `json:"time_type" validate:"required,oneof=days months years"`
	Duration     int    `json:"duration" validate:"required,min=1"`
}

func (r *UpdateLicenseRequest) ToV1Domain() *V1Domains.LicenseDomain {
	return &V1Domains.LicenseDomain{
		LicenseType: r.LicenseType,
	}
}

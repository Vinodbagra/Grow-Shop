package requests

import (
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
)

type UpdateLicenseRequest struct {
	LicenseType string `json:"license_type" validate:"required,oneof=FREE SILVER GOLD other"`
	ShopLimit   int    `json:"shop_limit" validate:"required"`
}

func (r *UpdateLicenseRequest) ToV1Domain() *V1Domains.LicenseDomain {
	return &V1Domains.LicenseDomain{
		LicenseType: r.LicenseType,
		ShopLimit:   r.ShopLimit,
	}
}

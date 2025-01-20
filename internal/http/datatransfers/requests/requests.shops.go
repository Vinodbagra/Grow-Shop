package requests

import V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"

// General Request
type ShopRequest struct {
	ShopName string `json:"name" validate:"required"`
}

// Mapping General Request to Domain User
func (shop ShopRequest) ToV1Domain() *V1Domains.ShopDomain {
	return &V1Domains.ShopDomain{
		ShopName: shop.ShopName,
	}
}

type UpdateShopRequest struct {
	MobileNo     string `json:"mobile_no"`
	Address      string `json:"address" validate:"max=255"`
	BusinessName string `json:"business_name" validate:"max=100"`
	Gender       string `json:"gender" validate:"oneof=male female other"`
	// Images       []string `json:"images" validate:"omitempty,dive,url"`
}

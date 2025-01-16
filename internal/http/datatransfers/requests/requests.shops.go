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

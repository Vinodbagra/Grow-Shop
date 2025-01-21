package responses

import (
	"time"

	"github.com/google/uuid"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
)

type ShopResponse struct {
	ShopName  string    `json:"shop_name"`
	UserID    uuid.UUID `json:"id"`
	RoleId    int       `json:"role_id"`
	Token     string    `json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *ShopResponse) ToV1Domain() V1Domains.ShopDomain {
	return V1Domains.ShopDomain{
		ShopName:  s.ShopName,
		UserID:    s.UserID,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func FromShopV1Domain(s V1Domains.ShopDomain) ShopResponse {
	return ShopResponse{
		ShopName:  s.ShopName,
		UserID:    s.UserID,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

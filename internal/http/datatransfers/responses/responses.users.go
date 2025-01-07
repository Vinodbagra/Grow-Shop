package responses

import (
	"time"

	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
)

type UserResponse struct {
	Id        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	RoleId    int        `json:"role_id"`
	Token     string     `json:"token,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *UserResponse) ToV1Domain() V1Domains.UserDomain {
	return V1Domains.UserDomain{
		UserID:        u.Id,
		UserName:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func FromV1Domain(u V1Domains.UserDomain) UserResponse {
	return UserResponse{
		Id:        u.UserID,
		Username:  u.UserName,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToResponseList(domains []V1Domains.UserDomain) []UserResponse {
	var result []UserResponse

	for _, val := range domains {
		result = append(result, FromV1Domain(val))
	}

	return result
}

package records

import (
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
)

func (u *Users) ToV1Domain() V1Domains.UserDomain {
	return V1Domains.UserDomain{
		UserID:    u.UserID,
		UserName:  u.UserName,
		Email:     u.Email,
		Password:  u.UserPassword,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func FromUsersV1Domain(u *V1Domains.UserDomain) Users {
	return Users{
		UserID:       u.UserID,
		UserName:     u.UserName,
		Email:        u.Email,
		UserPassword: u.Password,
		MobileNo:     u.MobileNo,
		Address:      u.Address,
		BusinessName: u.BusinessName,
		Gender:       u.Gender,
		Shops:        u.Shops,
		LicenseID:    u.LicenseID,
		Images:       u.Images,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func ToArrayOfUsersV1Domain(u *[]Users) []V1Domains.UserDomain {
	var result []V1Domains.UserDomain

	for _, val := range *u {
		result = append(result, val.ToV1Domain())
	}

	return result
}

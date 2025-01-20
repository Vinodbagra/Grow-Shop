package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/internal/datasources/caches"
	V1PostgresRepository "github.com/snykk/grow-shop/internal/datasources/repositories/postgres/v1"
	"github.com/snykk/grow-shop/pkg/helpers"
	"github.com/snykk/grow-shop/pkg/logger"
)

type shopservice struct {
	repo       V1PostgresRepository.ShopRepository
	redisCache caches.RedisCache
}

type ShopService interface {
	CreateShop(ctx context.Context, inDom *V1Domains.ShopDomain) (outDom V1Domains.ShopDomain, statusCode int, err error)
}

func NewShopService(repo V1PostgresRepository.ShopRepository, redisCache caches.RedisCache) ShopService {
	return &shopservice{
		repo:       repo,
		redisCache: redisCache,
	}
}

func (shop *shopservice) CreateShop(ctx context.Context, inDom *V1Domains.ShopDomain) (outDom V1Domains.ShopDomain, statusCode int, err error) {
	inDom.Password, err = helpers.GenerateHash(inDom.Password)
	if err != nil {
		return V1Domains.ShopDomain{}, http.StatusInternalServerError, err
	}

	inDom.CreatedAt = time.Now().In(constants.GMT7)
	fmt.Println(time.Now().In(constants.GMT7))
	err = shop.repo.CreateShop(ctx, inDom)
	if err != nil {
		return V1Domains.ShopDomain{}, http.StatusInternalServerError, err
	}

	outDom, err = shop.repo.GetByEmail(ctx, inDom)
	if err != nil {
		return V1Domains.ShopDomain{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusCreated, nil
}

func (u *shopservice) GetShopByID(ctx context.Context, shopID uuid.UUID) (outDom V1Domains.ShopDomain, statusCode int, err error) {
	user, err := u.repo.GetShopByID(ctx, &V1Domains.UserDomain{UserID: shopID})
	if err != nil {
		return V1Domains.ShopDomain{}, http.StatusNotFound, errors.New("user not found")
	}

	return user, http.StatusOK, nil
}

// create a function to update user data
func (u *shopservice) UpdateUserData(ctx context.Context, shopData *V1Domains.ShopDomain) (statusCode int, err error) {
	err = u.repo.UpdateShopData(ctx, shopData)
	if err != nil {
		logger.ErrorF("failed to update user data", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

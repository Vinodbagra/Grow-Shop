package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	V1services "github.com/snykk/grow-shop/internal/business/service/v1"
	"github.com/snykk/grow-shop/internal/datasources/caches"
	"github.com/snykk/grow-shop/internal/http/datatransfers/requests"
	"github.com/snykk/grow-shop/internal/http/datatransfers/responses"
	"github.com/snykk/grow-shop/pkg/validators"
)

type ShopHandler struct {
	service    V1services.ShopService
	redisCache caches.RedisCache
}

func NewShopHandler(service V1services.ShopService, redisCache caches.RedisCache) ShopHandler {
	return ShopHandler{
		service:    service,
		redisCache: redisCache,
	}
}

func (shopH ShopHandler) CreateShop(ctx *gin.Context) {
	var ShopCreateRequest requests.ShopRequest
	if err := ctx.ShouldBindJSON(&ShopCreateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(ShopCreateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	shopDomain := ShopCreateRequest.ToV1Domain()
	shopDomainn, statusCode, err := shopH.service.CreateShop(ctx.Request.Context(), shopDomain)
	//fmt.Println(shopDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "shop cration success", map[string]interface{}{
		"shop": responses.FromShopV1Domain(shopDomainn),
	})
}

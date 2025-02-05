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

func (c ShopHandler) GetShopData(ctx *gin.Context) {
	// Get userID from the context
	// , exists := ctx.Get("shopID")
	// if !exists {
	// 	NewErrorResponse(ctx, http.StatusUnauthorized, "shop ID not found in context")
	// 	return
	// }

	// Convert userID to the expected type
	// shopUUID, ok := shopID.(uuid.UUID)
	// if !ok {
	// 	NewErrorResponse(ctx, http.StatusInternalServerError, "invalid shop ID format")
	// 	return
	// }

	// Call the service to fetch user data by userID
	// ctxx := ctx.Request.Context()
	// shopDom, statusCode, err := c.service.GetShopByID(ctxx, shopUUID)
	// if err != nil {
	// 	NewErrorResponse(ctx, statusCode, err.Error())
	// 	return
	// }

	// Transform user domain data to a response format
	// shopResponse := responses.FromShopV1Domain(shopDom)

	// // Send success response
	// NewSuccessResponse(ctx, statusCode, "shop data fetched successfully", map[string]interface{}{
	// 	"shop": shopResponse,
	// })
}

// create a function for edit use data
func (shopH ShopHandler) UpdateShopData(ctx *gin.Context) {
	var ShopUpdateRequest requests.UpdateShopRequest
	if err := ctx.ShouldBindJSON(&ShopUpdateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(ShopUpdateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// shopDomain := ShopUpdateRequest.ToV1Domain()

	// // Extract user ID from the URL parameter or request context (adjust based on your setup)
	// shopID := ctx.Param("id")
	// if shopID == "" {
	// 	NewErrorResponse(ctx, http.StatusBadRequest, "user ID is required")
	// 	return
	// }

	// uuidShopID, err := uuid.Parse(shopID)
	// if err != nil {
	// 	NewErrorResponse(ctx, http.StatusBadRequest, "invalid shop ID format")
	// 	return
	// }

	// shopDomain.ShopID = uuidShopID

	// statusCode, err := shopH.service.UpdateShopData(ctx.Request.Context(), shopDomain)
	// if err != nil {
	// 	NewErrorResponse(ctx, statusCode, err.Error())
	// 	return
	// }

	// NewSuccessResponse(ctx, statusCode, "shop updated successfully", nil)
}

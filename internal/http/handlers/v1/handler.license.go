package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	V1services "github.com/snykk/grow-shop/internal/business/service/v1"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/internal/datasources/caches"
	"github.com/snykk/grow-shop/internal/http/datatransfers/requests"
	"github.com/snykk/grow-shop/pkg/validators"
)

type LicenseHandler struct {
	service    V1services.LicenseService
	redisCache caches.RedisCache
}

func NewLicenseHandler(service V1services.LicenseService, redisCache caches.RedisCache) LicenseHandler {
	return LicenseHandler{
		service:    service,
		redisCache: redisCache,
	}
}

func (licenseH LicenseHandler) UpdateLicenseData(ctx *gin.Context) {
	var licenseUpdateRequest requests.UpdateLicenseRequest
	if err := ctx.ShouldBindJSON(&licenseUpdateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(licenseUpdateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	licenseDomain := licenseUpdateRequest.ToV1Domain()

	// Extract user ID from the URL parameter or request context (adjust based on your setup)
	licenseID := ctx.Param("id")
	if licenseID == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "license ID is required")
		return
	}

	uuidLicenseID, err := uuid.Parse(licenseID)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid license ID format")
		return
	}

	licenseDomain.LicenseID = uuidLicenseID

	userID, exist := ctx.Get("userID")
	if !exist {
		NewErrorResponse(ctx, http.StatusUnauthorized, "user ID not found in context")
		return
	}
	userUUID, err := getUUIDFromContext(userID)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, "invalid user ID format")
		return
	}
	licenseDomain.UserID = userUUID
	statusCode, err := licenseH.service.UpdateLicense(ctx.Request.Context(), licenseDomain)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "license updated successfully", nil)
}

func getUUIDFromContext(value any) (uuid.UUID, error) {
	// Check if the value is of type uuid.UUID
	if userID, ok := value.(uuid.UUID); ok {
		return userID, nil
	}

	// If the value is a string, attempt to parse it as a UUID
	if userIDStr, ok := value.(string); ok {
		parsedUUID, err := uuid.Parse(userIDStr)
		if err != nil {
			return uuid.UUID{}, constants.ErrInvalidUUIDFormat
		}
		return parsedUUID, nil
	}

	return uuid.UUID{}, constants.ErrInvalidUUIDFormat
}

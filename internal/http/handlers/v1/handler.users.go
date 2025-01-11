package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	V1services "github.com/snykk/grow-shop/internal/business/service/v1"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/internal/datasources/caches"
	"github.com/snykk/grow-shop/internal/http/datatransfers/requests"
	"github.com/snykk/grow-shop/internal/http/datatransfers/responses"
	"github.com/snykk/grow-shop/pkg/logger"
	"github.com/snykk/grow-shop/pkg/validators"
)

type UserHandler struct {
	service    V1services.Userservice
	redisCache caches.RedisCache
}

func NewUserHandler(service V1services.Userservice, redisCache caches.RedisCache) UserHandler {
	return UserHandler{
		service:    service,
		redisCache: redisCache,
	}
}

func (userH UserHandler) Regis(ctx *gin.Context) {
	var UserRegisRequest requests.UserRequest
	if err := ctx.ShouldBindJSON(&UserRegisRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(UserRegisRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userDomain := UserRegisRequest.ToV1Domain()
	userDomainn, statusCode, err := userH.service.Store(ctx.Request.Context(), userDomain)
	fmt.Println(userDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "registration user success", map[string]interface{}{
		"user": responses.FromV1Domain(userDomainn),
	})
}

func (userH UserHandler) Login(ctx *gin.Context) {
	methodName := "Handler.Login"
	logger.InfoF("function name %s recieved the request to login", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, methodName)
	var UserLoginRequest requests.UserLoginRequest
	if err := ctx.ShouldBindJSON(&UserLoginRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(UserLoginRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, statusCode, err := userH.service.Login(ctx.Request.Context(), UserLoginRequest.ToV1Domain())
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "login success", map[string]interface{}{
		"token": token,
	})
}

func (userH UserHandler) ForgotPassword(ctx *gin.Context) {
	var userForgotPasswordRequest requests.UserForgotPasswordRequest

	if err := ctx.ShouldBindJSON(&userForgotPasswordRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(userForgotPasswordRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	statusCode, err := userH.service.ForgotPassword(ctx.Request.Context(), userForgotPasswordRequest.Email)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, fmt.Sprintf("reset link has been send to your email: %s", userForgotPasswordRequest.Email), nil)
}

func (userH UserHandler) VerifOTP(ctx *gin.Context) {
	var userOTP requests.UserVerifOTPRequest

	if err := ctx.ShouldBindJSON(&userOTP); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(userOTP); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	otpKey := fmt.Sprintf("user_otp:%s", userOTP.Email)
	otpRedis, err := userH.redisCache.Get(otpKey)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	statusCode, err := userH.service.VerifOTP(ctx.Request.Context(), userOTP.Email, userOTP.Code, otpRedis)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	statusCode, err = userH.service.ActivateUser(ctx.Request.Context(), userOTP.Email)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	go userH.redisCache.Del(otpKey)

	NewSuccessResponse(ctx, statusCode, "otp verification success", nil)
}

func (c UserHandler) GetUserData(ctx *gin.Context) {
	// Get userID from the context
	userID, exists := ctx.Get("userID")
	if !exists {
		NewErrorResponse(ctx, http.StatusUnauthorized, "user ID not found in context")
		return
	}

	// Convert userID to the expected type
	userUUID, ok := userID.(uuid.UUID)
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "invalid user ID format")
		return
	}

	// Call the service to fetch user data by userID
	ctxx := ctx.Request.Context()
	userDom, statusCode, err := c.service.GetUserByID(ctxx, userUUID)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	// Transform user domain data to a response format
	userResponse := responses.FromV1Domain(userDom)

	// Send success response
	NewSuccessResponse(ctx, statusCode, "user data fetched successfully", map[string]interface{}{
		"user": userResponse,
	})
}

// create a function for edit use data

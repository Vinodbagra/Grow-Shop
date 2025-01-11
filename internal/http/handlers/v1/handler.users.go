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

func (userH UserHandler) Register(ctx *gin.Context) {
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

func (userH UserHandler) ResetPassword(ctx *gin.Context) {
	var userResetPasswordRequest requests.UserResetPasswordRequest

	if err := ctx.ShouldBindJSON(& userResetPasswordRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads( userResetPasswordRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userDomain := userResetPasswordRequest.ToV1Domain()
	statusCode, err := userH.service.ResetPassword(ctx.Request.Context(), userDomain,  userResetPasswordRequest.ResetToken)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "password reset success", nil)
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
func (userH UserHandler) UpdateUserData(ctx *gin.Context) {
    var UserUpdateRequest requests.UpdateUserRequest
    if err := ctx.ShouldBindJSON(&UserUpdateRequest); err != nil {
        NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
        return
    }

    if err := validators.ValidatePayloads(UserUpdateRequest); err != nil {
        NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
        return
    }

    userDomain := UserUpdateRequest.ToV1Domain()

    // Extract user ID from the URL parameter or request context (adjust based on your setup)
    userID := ctx.Param("id")
    if userID == "" {
        NewErrorResponse(ctx, http.StatusBadRequest, "user ID is required")
        return
    }

    uuidUserID, err := uuid.Parse(userID)
    if err != nil {
        NewErrorResponse(ctx, http.StatusBadRequest, "invalid user ID format")
        return
    }

    userDomain.UserID = uuidUserID

    statusCode, err := userH.service.UpdateUserData(ctx.Request.Context(), userDomain)
    if err != nil {
        NewErrorResponse(ctx, statusCode, err.Error())
        return
    }

    NewSuccessResponse(ctx, statusCode, "user updated successfully", nil)
}

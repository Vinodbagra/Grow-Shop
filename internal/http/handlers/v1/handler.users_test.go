package v1_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	dgriJWT "github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// 	V1Domains "github.com/snykk/grow-shop/internal/business/domains/v1"
// 	V1services "github.com/snykk/grow-shop/internal/business/service/v1"
// 	"github.com/snykk/grow-shop/internal/config"
// 	"github.com/snykk/grow-shop/internal/constants"
// 	"github.com/snykk/grow-shop/internal/http/datatransfers/requests"
// 	V1Handlers "github.com/snykk/grow-shop/internal/http/handlers/v1"
// 	"github.com/snykk/grow-shop/internal/mocks"
// 	"github.com/snykk/grow-shop/pkg/helpers"
// 	"github.com/snykk/grow-shop/pkg/jwt"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// var (
// 	jwtServiceMock  *mocks.JWTService
// 	userRepoMock    *mocks.UserRepository
// 	userservice     V1services.Userservice
// 	userHandler     V1Handlers.UserHandler
// 	mailerOTPMock   *mocks.OTPMailer
// 	usersDataFromDB []V1Domains.UserDomain
// 	userDataFromDB  V1Domains.UserDomain
// 	redisMock       *mocks.RedisCache
// 	ristrettoMock   *mocks.RistrettoCache
// 	s               *gin.Engine
// )

// func setup(t *testing.T) {
// 	jwtServiceMock = mocks.NewJWTService(t)
// 	redisMock = mocks.NewRedisCache(t)
// 	mailerOTPMock = mocks.NewOTPMailer(t)
// 	ristrettoMock = mocks.NewRistrettoCache(t)
// 	userRepoMock = mocks.NewUserRepository(t)
// 	tokenRepoMock := mocks.NewTokenRepository(t)
// 	userservice = V1services.NewUserservice(userRepoMock, tokenRepoMock, mailerOTPMock)
// 	userHandler = V1Handlers.NewUserHandler(userservice, redisMock)

// 	usersDataFromDB = []V1Domains.UserDomain{
// 		{
// 			UserID:    "ddfcea5c-d919-4a8f-a631-4ace39337s3a",
// 			UserName:  "itsmepatrick",
// 			Email:     "najibfikri13@gmail.com",
// 			Password:  "23123sdf!",
// 			CreatedAt: time.Now(),
// 		},
// 		{
// 			UserID:    "wifff3jd-idhd-0sis-8dua-4fiefie37kfj",
// 			UserName:  "johny",
// 			Email:     "johny123@gmail.com",
// 			Password:  "23123sdf!",
// 			CreatedAt: time.Now(),
// 		},
// 	}

// 	userDataFromDB = V1Domains.UserDomain{
// 		UserID:    "fjskeie8-jfk8-qke0-sksj-ksjf89e8ehfu",
// 		UserName:  "itsmepatrick",
// 		Email:     "najibfikri13@gmail.com",
// 		Password:  "23123sdf!",
// 		CreatedAt: time.Now(),
// 	}

// 	// Create gin engine
// 	s = gin.Default()
// 	s.Use(lazyAuth)
// }

// func lazyAuth(ctx *gin.Context) {
// 	// hash
// 	pass, _ := helpers.GenerateHash(userDataFromDB.Password)
// 	// prepare claims
// 	jwtClaims := jwt.JwtCustomClaim{
// 		UserID:   userDataFromDB.UserID,
// 		IsAdmin:  false,
// 		Email:    userDataFromDB.Email,
// 		Password: pass,
// 		StandardClaims: dgriJWT.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * time.Duration(config.AppConfig.JWTExpired)).Unix(),
// 			Issuer:    userDataFromDB.UserName,
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 	}
// 	ctx.Set(constants.CtxAuthenticatedUserKey, jwtClaims)
// }

// func TestRegis(t *testing.T) {
// 	setup(t)
// 	// Define route
// 	s.POST(constants.EndpointV1+"/auth/regis", userHandler.Regis)
// 	t.Run("When Success Regis", func(t *testing.T) {
// 		req := requests.UserRequest{
// 			Email:    "najibfikri13@gmail.com",
// 			Password: "23123sdf!",
// 		}
// 		reqBody, _ := json.Marshal(req)

// 		userRepoMock.Mock.On("Store", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(nil).Once()
// 		userRepoMock.Mock.On("GetByEmail", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(userDataFromDB, nil).Once()

// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodPost, constants.EndpointV1+"/auth/regis", bytes.NewReader(reqBody))

// 		r.Header.Set("Content-Type", "application/json")

// 		// Perform request
// 		s.ServeHTTP(w, r)

// 		body := w.Body.String()

// 		// Assertions
// 		// Assert status code
// 		assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
// 		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 		assert.Contains(t, body, "registration user success")
// 	})
// 	t.Run("When Failure", func(t *testing.T) {
// 		t.Run("When Request is Empty", func(t *testing.T) {
// 			req := requests.UserRequest{}
// 			reqBody, _ := json.Marshal(req)

// 			w := httptest.NewRecorder()
// 			r := httptest.NewRequest(http.MethodPost, constants.EndpointV1+"/auth/regis", bytes.NewReader(reqBody))

// 			r.Header.Set("Content-Type", "application/json")

// 			// Perform request
// 			s.ServeHTTP(w, r)

// 			body := w.Body.String()

// 			// Assertions
// 			// Assert status code
// 			assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
// 			assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 			assert.Contains(t, body, "required")
// 		})
// 	})
// }

// func TestForgotPassword(t *testing.T) {
// 	setup(t)
// 	// Define route
// 	s.POST(constants.EndpointV1+"/auth/send-otp", userHandler.ForgotPassword)
// 	t.Run("Test 1 | Success Send OTP", func(t *testing.T) {
// 		req := requests.UserForgotPasswordRequest{
// 			Email: "najibfikri13@gmail.com",
// 		}
// 		reqBody, _ := json.Marshal(req)

// 		userRepoMock.Mock.On("GetByEmail", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(userDataFromDB, nil).Once()
// 		mailerOTPMock.On("ForgotPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil).Once()
// 		redisMock.On("Set", mock.AnythingOfType("string"), mock.Anything).Return(nil).Once()

// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodPost, constants.EndpointV1+"/auth/send-otp", bytes.NewReader(reqBody))

// 		r.Header.Set("Content-Type", "application/json")

// 		// Perform request
// 		s.ServeHTTP(w, r)

// 		body := w.Body.String()

// 		// Assertions
// 		// Assert status code
// 		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
// 		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 		assert.Contains(t, body, "otp code has been send")
// 	})
// 	t.Run("Test 3 | Payloads is Empty", func(t *testing.T) {
// 		req := requests.UserForgotPasswordRequest{}
// 		reqBody, _ := json.Marshal(req)

// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodPost, constants.EndpointV1+"/auth/send-otp", bytes.NewReader(reqBody))

// 		r.Header.Set("Content-Type", "application/json")

// 		// Perform request
// 		s.ServeHTTP(w, r)

// 		// Assertions
// 		// Assert status code
// 		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
// 		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 	})
// 	t.Run("Test 3 | When Failure Send OTP", func(t *testing.T) {
// 		req := requests.UserForgotPasswordRequest{
// 			Email: "najibfikri13@gmail.com",
// 		}
// 		reqBody, _ := json.Marshal(req)

// 		userRepoMock.Mock.On("GetByEmail", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(userDataFromDB, nil).Once()
// 		mailerOTPMock.On("ForgotPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(constants.ErrUnexpected).Once()

// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodPost, constants.EndpointV1+"/auth/send-otp", bytes.NewReader(reqBody))

// 		r.Header.Set("Content-Type", "application/json")

// 		// Perform request
// 		s.ServeHTTP(w, r)

// 		// Assertions
// 		// Assert status code
// 		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
// 		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 	})
// }

// func TestResetPassword(t *testing.T) {
// 	setup(t)
// 	// Define route
// 	s.POST(constants.EndpointV1+"/auth/verif-otp", userHandler.ResetPassword)
// 	t.Run("Test 1 | Success Verify OTP", func(t *testing.T) {
// 		req := requests.UserResetPasswordRequest{
// 			Email: "najibfikri13@gmail.com",
// 			Code:  "112233",
// 		}
// 		reqBody, _ := json.Marshal(req)

// 		redisMock.Mock.On("Get", mock.AnythingOfType("string")).Return("112233", nil)
// 		redisMock.On("Del", mock.AnythingOfType("string")).Return(nil).Once()
// 		ristrettoMock.On("Del", mock.AnythingOfType("string")).Once()

// 		userRepoMock.Mock.On("GetByEmail", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(userDataFromDB, nil).Once()
// 		userRepoMock.Mock.On("GetByEmail", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(userDataFromDB, nil).Once()
// 		userRepoMock.Mock.On("ChangeActiveUser", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(nil).Once()

// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodPost, constants.EndpointV1+"/auth/verif-otp", bytes.NewReader(reqBody))

// 		r.Header.Set("Content-Type", "application/json")

// 		// Perform request
// 		s.ServeHTTP(w, r)

// 		body := w.Body.String()

// 		// Assertions
// 		// Assert status code
// 		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
// 		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 		assert.Contains(t, body, "otp verification success")
// 	})
// 	t.Run("Test 2 | Payloads is Empty", func(t *testing.T) {
// 		req := requests.UserResetPasswordRequest{}
// 		reqBody, _ := json.Marshal(req)

// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodPost, constants.EndpointV1+"/auth/verif-otp", bytes.NewReader(reqBody))

// 		r.Header.Set("Content-Type", "application/json")

// 		// Perform request
// 		s.ServeHTTP(w, r)

// 		// Assertions
// 		// Assert status code
// 		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
// 		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 	})
// 	t.Run("Test 1 | Invalid OTP Code", func(t *testing.T) {
// 		req := requests.UserResetPasswordRequest{
// 			Email: "najibfikri13@gmail.com",
// 			Code:  "999999",
// 		}
// 		reqBody, _ := json.Marshal(req)

// 		redisMock.Mock.On("Get", mock.AnythingOfType("string")).Return("112233", nil)

// 		userRepoMock.Mock.On("GetByEmail", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(userDataFromDB, nil).Once()

// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodPost, constants.EndpointV1+"/auth/verif-otp", bytes.NewReader(reqBody))

// 		r.Header.Set("Content-Type", "application/json")

// 		// Perform request
// 		s.ServeHTTP(w, r)

// 		body := w.Body.String()

// 		// Assertions
// 		// Assert status code
// 		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
// 		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 		assert.Contains(t, body, "invalid otp code")
// 	})
// }

// func TestLogin(t *testing.T) {
// 	setup(t)
// 	// Define route
// 	s.POST(constants.EndpointV1+"/auth/login", userHandler.Login)
// 	t.Run("Test 1 | Success Login", func(t *testing.T) {
// 		// hash password field
// 		var err error
// 		userDataFromDB.Password, err = helpers.GenerateHash(userDataFromDB.Password)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		// make account activated
// 		req := requests.UserLoginRequest{
// 			Email:    "patrick@gmail.com",
// 			Password: "23123sdf!",
// 		}
// 		reqBody, _ := json.Marshal(req)

// 		userRepoMock.Mock.On("GetByEmail", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(userDataFromDB, nil).Once()
// 		jwtServiceMock.Mock.On("GenerateToken", mock.AnythingOfType("string"), mock.AnythingOfType("bool"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("eyBlablablabla", nil).Once()

// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodPost, constants.EndpointV1+"/auth/login", bytes.NewReader(reqBody))

// 		r.Header.Set("Content-Type", "application/json")

// 		// Perform request
// 		s.ServeHTTP(w, r)

// 		body := w.Body.String()

// 		// Assertions
// 		// Assert status code
// 		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
// 		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 		assert.Contains(t, body, "login success")
// 		assert.Contains(t, body, "ey")
// 	})
// 	t.Run("Test 2 | User is Not Exists", func(t *testing.T) {
// 		req := requests.UserLoginRequest{
// 			Email:    "patrick312@gmail.com",
// 			Password: "23123sdf!",
// 		}
// 		reqBody, _ := json.Marshal(req)

// 		userRepoMock.Mock.On("GetByEmail", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(V1Domains.UserDomain{}, constants.ErrUserNotFound).Once()

// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodPost, constants.EndpointV1+"/auth/login", bytes.NewReader(reqBody))

// 		r.Header.Set("Content-Type", "application/json")

// 		// Perform request
// 		s.ServeHTTP(w, r)

// 		body := w.Body.String()

// 		// Assertions
// 		// Assert status code
// 		assert.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
// 		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 		assert.Contains(t, body, "invalid email or password")
// 	})
// }

// func TestGetUserData(t *testing.T) {
// 	setup(t)
// 	// Define route
// 	s.GET("/users/me", userHandler.GetUserData)

// 	authenticatedUserEmail := userDataFromDB.Email
// 	t.Run("Test 1 | Success Fetched User Data", func(t *testing.T) {
// 		userRepoMock.Mock.On("GetByEmail", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(userDataFromDB, nil).Once()
// 		ristrettoMock.Mock.On("Get", fmt.Sprintf("user/%s", authenticatedUserEmail)).Return(nil).Once()
// 		ristrettoMock.Mock.On("Set", fmt.Sprintf("user/%s", authenticatedUserEmail), mock.Anything).Once()

// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodGet, "/users/me", nil)

// 		r.Header.Set("Content-Type", "application/json")

// 		// Perform request
// 		s.ServeHTTP(w, r)

// 		// parsing json to raw text
// 		body := w.Body.String()

// 		// Assertions
// 		// Assert status code
// 		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
// 		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 		assert.Contains(t, body, "user data fetched successfully")
// 	})

// 	t.Run("Test 2 | Failed to fetch User Data", func(t *testing.T) {
// 		userRepoMock.Mock.On("GetByEmail", mock.Anything, mock.AnythingOfType("*v1.UserDomain")).Return(V1Domains.UserDomain{}, constants.ErrUnexpected).Once()
// 		ristrettoMock.Mock.On("Get", fmt.Sprintf("user/%s", authenticatedUserEmail)).Return(nil).Once()

// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodGet, "/users/me", nil)

// 		r.Header.Set("Content-Type", "application/json")

// 		// Perform request
// 		s.ServeHTTP(w, r)

// 		// Assertions
// 		// Assert status code
// 		assert.NotEqual(t, http.StatusOK, w.Result().StatusCode)
// 		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
// 	})
// }

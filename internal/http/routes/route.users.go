package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	V1service "github.com/snykk/grow-shop/internal/business/service/v1"
	"github.com/snykk/grow-shop/internal/datasources/caches"
	V1PostgresRepository "github.com/snykk/grow-shop/internal/datasources/repositories/postgres/v1"
	V1Handler "github.com/snykk/grow-shop/internal/http/handlers/v1"
	"github.com/snykk/grow-shop/pkg/mailer"
)

type usersRoutes struct {
	V1Handler      V1Handler.UserHandler
	router         *gin.RouterGroup
	db             *sqlx.DB
	authMiddleware gin.HandlerFunc
}

func NewUsersRoute(router *gin.RouterGroup, db *sqlx.DB, redisCache caches.RedisCache, authMiddleware gin.HandlerFunc, mailer mailer.OTPMailer) *usersRoutes {
	V1UserRepository := V1PostgresRepository.NewUserRepository(db)
	V1TokenRepository := V1PostgresRepository.NewTokenRepository(db)
	V1Userservice := V1service.NewUserservice(V1UserRepository,V1TokenRepository, mailer)
	V1UserHandler := V1Handler.NewUserHandler(V1Userservice, redisCache)

	return &usersRoutes{V1Handler: V1UserHandler, router: router, db: db, authMiddleware: authMiddleware}
}

func (r *usersRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// auth
		V1AuhtRoute := V1Route.Group("/auth")
		V1AuhtRoute.POST("/regis", r.V1Handler.Regis)
		V1AuhtRoute.POST("/login", r.V1Handler.Login)
		// V1AuhtRoute.POST("/send-otp", r.V1Handler.SendOTP)
		// V1AuhtRoute.POST("/verif-otp", r.V1Handler.VerifOTP)

		// users
		userRoute := V1Route.Group("/users")
		userRoute.Use(r.authMiddleware)
		{
			userRoute.GET("/", r.V1Handler.GetUserData)
			// ...
		}
	}

}

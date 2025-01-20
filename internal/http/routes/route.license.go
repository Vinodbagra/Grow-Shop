package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	V1service "github.com/snykk/grow-shop/internal/business/service/v1"
	"github.com/snykk/grow-shop/internal/datasources/caches"
	V1PostgresRepository "github.com/snykk/grow-shop/internal/datasources/repositories/postgres/v1"
	V1Handler "github.com/snykk/grow-shop/internal/http/handlers/v1"
)

type licenseRoutes struct {
	V1Handler      V1Handler.LicenseHandler
	router         *gin.RouterGroup
	db             *sqlx.DB
	authMiddleware gin.HandlerFunc
}

func NewLicenseRoute(router *gin.RouterGroup, db *sqlx.DB, redisCache caches.RedisCache, authMiddleware gin.HandlerFunc) *licenseRoutes {
	V1LicenseRepository := V1PostgresRepository.NewLicenseRepository(db)
	V1Licenseservice := V1service.NewLicenseservice(V1LicenseRepository)
	V1LicenseHandler := V1Handler.NewLicenseHandler(V1Licenseservice, redisCache)

	return &licenseRoutes{V1Handler: V1LicenseHandler, router: router, db: db, authMiddleware: authMiddleware}
}

func (r *licenseRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// license
		licenseRoute := V1Route.Group("/licenses")
		licenseRoute.Use(r.authMiddleware)
		{
			// userRoute.GET("/", r.V1Handler.GetUserData)
			licenseRoute.PUT(":id", r.V1Handler.UpdateLicenseData)
			// ...
			// create a put api for updating user data
		}
	}

}

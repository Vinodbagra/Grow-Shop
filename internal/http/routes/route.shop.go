package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	V1service "github.com/snykk/grow-shop/internal/business/service/v1"
	"github.com/snykk/grow-shop/internal/datasources/caches"
	V1PostgresRepository "github.com/snykk/grow-shop/internal/datasources/repositories/postgres/v1"
	V1Handler "github.com/snykk/grow-shop/internal/http/handlers/v1"
)

// create a class with name shoproute
// function to create new object of this shop class with naem NewShopRoute
// create route function for this class

type shopsRoutes struct {
	V1Handler      V1Handler.ShopHandler // have to made shop handler after created in the handler
	router         *gin.RouterGroup
	db             *sqlx.DB
	authMiddleware gin.HandlerFunc
}

func NewShopsRoute(router *gin.RouterGroup, db *sqlx.DB, redisCache caches.RedisCache, authMiddleware gin.HandlerFunc) *shopsRoutes {
	V1ShopRepository := V1PostgresRepository.NewShopRepository(db)
	V1ShopService := V1service.NewShopService(V1ShopRepository, redisCache)
	V1ShopHandler := V1Handler.NewShopHandler(V1ShopService, redisCache)

	return &shopsRoutes{V1Handler: V1ShopHandler, router: router, db: db, authMiddleware: authMiddleware}
}


func (r *shopsRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// shops
		// POST --> to create new 
		// PUT --> to update existing data
		// GET --> to get existing data
		shopRoute := V1Route.Group("/shops")
		shopRoute.Use(r.authMiddleware)
		{
			//userRoute.GET("/", r.V1Handler.GetUserData)
			//userRoute.PUT(":id", r.V1Handler.UpdateUserData)
			// ...
			// create a put api for updating user data
			shopRoute.POST("/", r.V1Handler.CreateShop)
			//shopRoute.GET("/get", r.V1Handler.GetShop)
			// handler --> service --> repo
		}
	}

}
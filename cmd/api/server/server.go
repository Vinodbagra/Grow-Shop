package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/snykk/grow-shop/internal/config"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/internal/datasources/caches"
	"github.com/snykk/grow-shop/internal/http/middlewares"
	"github.com/snykk/grow-shop/internal/http/routes"
	"github.com/snykk/grow-shop/internal/utils"
	"github.com/snykk/grow-shop/pkg/logger"
	"github.com/snykk/grow-shop/pkg/mailer"
)

type App struct {
	HttpServer *http.Server
}

func NewApp() (*App, error) {
	// setup databases
	conn, err := utils.SetupPostgresConnection()
	if err != nil {
		return nil, err
	}

	// setup router
	router := setupRouter()

	// jwt service
	// jwtService := jwt.NewJWTService(config.AppConfig.JWTSecret, config.AppConfig.JWTIssuer, config.AppConfig.JWTExpired)

	// cache
	redisCache := caches.NewRedisCache(config.AppConfig.REDISHost, 0, config.AppConfig.REDISPassword, time.Duration(config.AppConfig.REDISExpired))

	// mailer
	mailerService := mailer.NewOTPMailer(config.AppConfig.OTPEmail, config.AppConfig.OTPPassword)

	// user middleware
	// user with valid basic token can access endpoint
	authMiddleware := middlewares.NewAuthMiddleware(conn, false)

	// admin middleware
	// only user with valid admin token can access endpoint
	_ = middlewares.NewAuthMiddleware(conn,true)

	// API Routes
	api := router.Group("api")
	api.GET("/", routes.RootHandler)
	routes.NewUsersRoute(api, conn, redisCache, authMiddleware, mailerService).Routes()
	routes.NewLicenseRoute(api, conn, redisCache, authMiddleware).Routes()

	// we can add web pages if needed
	// web := router.Group("web")
	// ...

	// setup http server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &App{
		HttpServer: server,
	}, nil
}

func (a *App) Run() (err error) {
	// Gracefull Shutdown
	go func() {
		logger.InfoF("success to listen and serve on :%d", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, config.AppConfig.Port)
		if err := a.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// make blocking channel and waiting for a signal
	<-quit
	logger.Info("shutdown server ...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.HttpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	logger.Info("timeout of 5 seconds.", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	logger.Info("server exiting", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	return
}

func setupRouter() *gin.Engine {
	// set the runtime mode
	var mode = gin.ReleaseMode
	if config.AppConfig.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	// create a new router instance
	router := gin.New()

	// set up middlewares
	router.Use(middlewares.CORSMiddleware())
	router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))
	router.Use(gin.Recovery())

	return router
}

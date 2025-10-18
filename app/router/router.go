package router

import (
	"pismo/config"
	"pismo/consts"
	"pismo/internal/handler"
	"pismo/internal/middleware"
	"pismo/internal/repo"
	"pismo/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Capsule struct {
	DB         *gorm.DB
	Repo       *repo.Repo
	Usecase    *usecase.Usecase
	Handler    *handler.Handler
	Middleware middleware.Middleware
	Config     *config.Config
}

func PrepareRouter(capsule *Capsule) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health"),
		gin.Recovery(),
	)

	config := config.GetConfig()
	if config.Environment != consts.PRODUCTION {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	middleware := capsule.Middleware

	router.Use(middleware.CORS())
	router.Use(middleware.Trace())

	api := router.Group("api")
	v1 := api.Group("v1")

	capsule.HealthRoutes(v1)
	capsule.AccountRoutes(v1)
	capsule.TransactionRoutes(v1)
	capsule.UserRoutes(v1)

	return router
}

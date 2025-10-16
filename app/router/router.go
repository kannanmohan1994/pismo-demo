package router

import (
	"xm/config"
	"xm/consts"
	"xm/internal/handler"
	"xm/internal/middleware"
	"xm/internal/repo"
	"xm/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		router.Static("/swaggerui/", "./swagger-ui")
	}

	middleware := capsule.Middleware

	router.Use(middleware.CORS())
	router.Use(middleware.Trace())

	v1 := router.Group("v1")

	capsule.HealthRoutes(v1)
	capsule.CompanyRoutes(v1)
	capsule.UserRoutes(v1)

	return router
}

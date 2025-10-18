package middleware

import (
	"pismo/config"
	"pismo/logger"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Authenticate() gin.HandlerFunc
	CORS() gin.HandlerFunc
	Trace(skippedURLs ...string) gin.HandlerFunc
}

type middleware struct {
	logger logger.Log
	config config.Config
}

func InitMiddleware(cfg config.Config, logger logger.Log) Middleware {
	return &middleware{
		config: cfg,
		logger: logger,
	}
}

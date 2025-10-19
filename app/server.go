package app

import (
	"pismo/app/router"
	"pismo/config"
	"pismo/consts"
	database "pismo/db"
	"pismo/internal/handler"
	"pismo/internal/middleware"
	"pismo/internal/repo"
	"pismo/internal/usecase"
	"pismo/internal/validation"
	"pismo/logger"
	"time"
)

func Start() {
	config.Setup()
	config := config.GetConfig()

	db, err := database.PrepareDatabase()
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger(config)
	tokenFunc := middleware.NewUserToken(consts.SigningMethodHS256, config.JWTSecretKey, time.Duration(config.AccessTokenExpiryDurationSeconds*int(time.Second)))

	mwr := middleware.InitMiddleware(*config, logger)
	repo := repo.InitRepo(db, logger)
	uc := usecase.Init(repo, tokenFunc, logger)
	validation := validation.Init(uc)
	hndlr := handler.Init(validation, logger)

	router := router.PrepareRouter(&router.Capsule{
		DB:         db,
		Repo:       repo,
		Usecase:    uc,
		Handler:    hndlr,
		Config:     config,
		Middleware: mwr,
	})

	logger.Infof("server running at port %s", config.ServerPort)
	err = router.Run(":" + config.ServerPort)
	if err != nil {
		logger.Infof("error running server - %s", err.Error())
	}
}

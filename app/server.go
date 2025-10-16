package app

import (
	"fmt"
	"xm/app/pubsub"
	"xm/app/router"
	"xm/config"
	database "xm/db"
	"xm/internal/handler"
	"xm/internal/middleware"
	kafkaNotifier "xm/internal/notifiers/kafka"
	"xm/internal/repo"
	"xm/internal/usecase"
	"xm/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Start() {
	config.Setup()
	config := config.GetConfig()

	db, err := database.PrepareDatabase()
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger(config)

	publisher, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", config.KafkaHost, config.KafkaPort),
	})
	if err != nil {
		logger.Errorf("failed to create producer: %s\n", err)
		panic(err)
	}

	defer publisher.Close()

	notifier := kafkaNotifier.InitKafka(publisher)

	middleware := middleware.InitMiddleware(*config, logger)
	dom := repo.InitRepo(db, logger)
	uc := usecase.Init(dom, middleware, logger, notifier)
	hndlr := handler.Init(uc, logger)

	router := router.PrepareRouter(&router.Capsule{
		DB:         db,
		Repo:       dom,
		Usecase:    uc,
		Handler:    hndlr,
		Config:     config,
		Middleware: middleware,
	})

	pubsub.PreparePubSub(&pubsub.CapsulePubSub{
		Logger:  logger,
		Config:  config,
		Usecase: uc,
	})

	logger.Infof("server running at port %s", config.ServerPort)
	err = router.Run(":" + config.ServerPort)
	if err != nil {
		logger.Infof("error running server - %s", err.Error())
	}
}

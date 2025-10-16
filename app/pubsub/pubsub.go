package pubsub

import (
	"fmt"
	"sync"
	"xm/config"
	"xm/consts"
	"xm/internal/usecase"
	"xm/logger"
)

type CapsulePubSub struct {
	Logger  logger.Log
	Config  *config.Config
	Usecase *usecase.Usecase
}

var (
	once sync.Once
)

func PreparePubSub(capsule *CapsulePubSub) {
	once.Do(func() {
		capsule.Serve()
	})
}

func (c *CapsulePubSub) Serve() {
	go c.RunConsumer(fmt.Sprintf("%s:%s", c.Config.KafkaHost, c.Config.KafkaPort), consts.KafkaGroupId, []string{
		consts.NotifyEventCompleted,
	})
}
